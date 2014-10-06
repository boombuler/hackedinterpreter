package runtime

import (
	"../token"
)

func NewStatements(s1, s2 *Callable) *Callable {
	return newCallable(nil, func(c *Context) (Value, error) {
		if c.forceExit() {
			return c.result, c.err
		}
		v1, err := s1.Call(c)
		if err != nil {
			return v1, err
		}
		if c.forceExit() {
			return c.result, c.err
		}
		return s2.Call(c)
	}, s1, s2)
}

func NewWhileLoop(cond, code *Callable, p *token.Token) *Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		var res Value = int(0)
		for {
			if c.forceExit() {
				return c.result, c.err
			}
			bv, err := callBool(cond, c)
			if err != nil {
				return bv, err
			}
			if !bv {
				break
			}
			if c.forceExit() {
				return c.result, c.err
			}
			res, err = code.Call(c)
			if err != nil {
				return res, err
			}
		}
		return res, nil
	}, cond, code)
}

func NewIfThenElse(cond, ifCode, elseCode *Callable, p *token.Token) *Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		if c.forceExit() {
			return c.result, c.err
		}
		bv, err := callBool(cond, c)
		if err != nil {
			return bv, err
		}
		if c.forceExit() {
			return c.result, c.err
		}
		if bv {
			return ifCode.Call(c)
		} else if elseCode != nil {
			return elseCode.Call(c)
		}
		return int(0), nil
	}, cond, ifCode, elseCode)
}

func NewForEach(vn string, lst, code *Callable, p *token.Token) *Callable {
	callable := newCallable(p, func(c *Context) (Value, error) {
		list, err := callList(lst, c)
		if err != nil {
			return list, err
		}
		var res Value = int(0)
		for _, val := range list.content {
			if c.forceExit() {
				return c.result, c.err
			}
			c.variables[vn] = val
			res, err = code.Call(c)
			if err != nil {
				return res, err
			}
		}
		return res, nil
	}, lst, code)
	callable.info = &setVarInfo{vn}
	return callable
}
