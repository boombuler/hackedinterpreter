package runtime

import (
	"../token"
)

func NewStatements(s1, s2 Callable) Callable {
	return newCallable(nil, func(c *Context) (Value, error) {
		if c.forceExit() {
			return c.result, c.err
		}
		v1, err := c.Call(s1)
		if err != nil {
			return v1, err
		}
		if c.forceExit() {
			return c.result, c.err
		}
		return c.Call(s2)
	}, s1, s2)
}

func NewWhileLoop(cond, code Callable, p *token.Token) Callable {
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
			res, err = c.Call(code)
			if err != nil {
				return res, err
			}
		}
		return res, nil
	}, cond, code)
}

func NewIfThenElse(cond, ifCode, elseCode Callable, p *token.Token) Callable {
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
			return c.Call(ifCode)
		} else if elseCode != nil {
			return c.Call(elseCode)
		}
		return int(0), nil
	}, cond, ifCode, elseCode)
}

func NewForEach(vn string, lst, code Callable, p *token.Token) Callable {
	return newCallable(p, func(c *Context) (Value, error) {
		list, err := callList(lst, c)
		if err != nil {
			return list, err
		}
		var res Value = int(0)
		for _, val := range list.content {
			if c.forceExit() {
				return c.result, c.err
			}
			c.SetVariable(vn, val)
			res, err = c.Call(code)
			if err != nil {
				return res, err
			}
		}
		return res, nil
	}, lst, code)
}
