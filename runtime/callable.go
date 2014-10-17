package runtime

import (
	"../token"
	"bytes"
	"fmt"
)

type CallableMetadata struct {
	*token.Token
	Children []Callable
}

type Callable interface {
	Meta() *CallableMetadata
	invoke(c *Context) (Value, error)
}

type callable struct {
	meta *CallableMetadata
	fn   func(c *Context) (Value, error)
}

func newCallable(token *token.Token, fn func(c *Context) (Value, error), children ...Callable) *callable {
	meta := &CallableMetadata{token, children}
	return &callable{meta, fn}
}

func (cf *callable) Meta() *CallableMetadata {
	return cf.meta
}

func (cf *callable) invoke(c *Context) (Value, error) {
	return cf.fn(c)
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
