package runtime

import (
	"../token"
	"bytes"
	"fmt"
)

type Callable struct {
	*token.Token
	fn func(c *Context) (Value, error)
}

func newCallable(token *token.Token, fn func(c *Context) (Value, error)) *Callable {
	return &Callable{token, fn}
}

func (cf *Callable) Call(c *Context) (Value, error) {
	if c.err != nil || c.result != nil {
		return c.result, c.err
	}
	val, err := c.exec(cf)
	if err != nil {
		re, ok := err.(*RuntimeError)
		if ok {
			if cf.Token != nil {
				re.CallStack = append(re.CallStack, &cf.Pos)
			}
			return nil, re
		} else {
			return nil, &RuntimeError{
				Message:   err.Error(),
				Position:  &cf.Pos,
				CallStack: make([]*token.Pos, 0),
			}
		}
	}
	return val, nil
}

type RuntimeError struct {
	Message   string
	Position  *token.Pos
	CallStack []*token.Pos
}

func (re *RuntimeError) Error() string {
	w := new(bytes.Buffer)
	fmt.Fprint(w, "runtime error:\n")
	fmt.Fprintf(w, "%s at Line %d / Column %d\n", re.Message, re.Position.Line, re.Position.Column)
	if len(re.CallStack) > 0 {
		fmt.Fprint(w, "stack trace:\n")
		for _, p := range re.CallStack {
			fmt.Fprintf(w, "  Line %d / Column %d\n", p.Line, p.Column)
		}
	}
	return w.String()
}
