package runtime

func NewStatements(s1, s2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
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
	})
}

func NewWhileLoop(cond, code Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		var res Value = Int(0)
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
	})
}

func NewIfThenElse(cond, ifCode, elseCode Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
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
		return Int(0), nil
	})
}

func NewForEach(vn string, lst, code Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		list, err := callList(lst, c)
		if err != nil {
			return list, err
		}
		var res Value = Int(0)
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
	})
}
