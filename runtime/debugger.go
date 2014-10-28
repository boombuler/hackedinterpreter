package runtime

import (
	"../token"
	"errors"
)

type DebuggerExecutionMode byte

const (
	Continue DebuggerExecutionMode = iota
	Step
	eval
)

type BreakEvent struct {
	Continue chan<- DebuggerExecutionMode
	Token    *token.Token
}

type Debugger struct {
	curCtx      *Context
	mainCtx     *Context
	curCallable Callable

	mode DebuggerExecutionMode

	breakEventChan chan<- *BreakEvent
	memReadBPs     map[string]bool
	memWriteBPs    map[string]bool
	codeBPs        map[*token.Token]bool // change bool to something else if conditional BPs should be impl.
}

func AttachDebugger(c *Context, be chan<- *BreakEvent) (*Debugger, error) {
	if c == nil || c.parentContext != nil {
		return nil, errors.New("Debugger can only be attached to main context")
	}
	d := &Debugger{
		mainCtx:        c,
		breakEventChan: be,
		mode:           Step,
		codeBPs:        make(map[*token.Token]bool),
		memReadBPs:     make(map[string]bool),
		memWriteBPs:    make(map[string]bool),
	}
	c.executor = d
	return d, nil
}

func (d *Debugger) ToggleCodeBreakPoint(t *token.Token) {
	_, ok := d.codeBPs[t]
	if ok {
		delete(d.codeBPs, t)
	} else {
		d.codeBPs[t] = true
	}
}

func (d *Debugger) IsCodeBreakPoint(t *token.Token) bool {
	_, ok := d.codeBPs[t]
	return ok
}

func (d *Debugger) ToggleMemReadBreakPoint(varName string) {
	_, ok := d.memReadBPs[varName]
	if ok {
		delete(d.memReadBPs, varName)
	} else {
		d.memReadBPs[varName] = true
	}
}

func (d *Debugger) ToggleMemWriteBreakPoint(varName string) {
	_, ok := d.memWriteBPs[varName]
	if ok {
		delete(d.memWriteBPs, varName)
	} else {
		d.memWriteBPs[varName] = true
	}
}

func (d *Debugger) IsMemBreakPoint(varName string) bool {
	_, ok := d.memReadBPs[varName]
	if ok {
		return true
	}
	_, ok = d.memWriteBPs[varName]
	return ok
}

func (d *Debugger) GetVars() map[string]Value {
	if d.curCtx != nil {
		return d.curCtx.variables
	}
	return map[string]Value{}
}

func (d *Debugger) Eval(c Callable) {
	orgMode := d.mode
	d.mode = eval
	defer func() { d.mode = orgMode }()
	d.curCtx.Call(c)
}

func (d *Debugger) getVar(vn string, ctx *Context) Value {
	v, ok := ctx.variables[vn]
	if !ok {
		v = 0
	}
	if d.mode != eval && d.memReadBPs[vn] {
		d.mode = Step
	}
	return v
}
func (d *Debugger) setVar(vn string, value Value, ctx *Context) {
	ctx.variables[vn] = value
	if d.mode != eval && d.memWriteBPs[vn] {
		d.mode = Step
	}
}

func (d *Debugger) exec(c Callable, ctx *Context) (Value, error) {
	meta := c.Meta()
	if meta != nil && meta.Token != nil && d.mode != eval {
		d.curCtx = ctx
		d.curCallable = c

		// Check here for BreakPoint etc
		if d.mode == Step || d.codeBPs[meta.Token] {
			// We reached a BP...
			if d.breakEventChan != nil {
				res := make(chan DebuggerExecutionMode)
				d.breakEventChan <- &BreakEvent{res, meta.Token}
				d.mode = <-res
			}
		}
	}
	return c.invoke(ctx)
}
