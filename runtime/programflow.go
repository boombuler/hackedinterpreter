package runtime

func MakeStatements(s1, s2 Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		v1, err := s1.Call(c)
		if err != nil {
			return v1, err
		}
		return s2.Call(c)
	})
}

func MakeWhileLoop(cond, code Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		var res Value = Int(0)
		for {
			bv, err := callBool(cond, c)
			if err != nil {
				return bv, err
			}
			if !bv {
				break
			}
			res, err = code.Call(c)
			if err != nil {
				return res, err
			}
		}
		return res, nil
	})
}

func MakeIfThenElse(cond, ifCode, elseCode Callable) Callable {
	return callableFunc(func(c *Context) (Value, error) {
		bv, err := callBool(cond, c)
		if err != nil {
			return bv, err
		}
		if bv {
			return ifCode.Call(c)
		} else if elseCode != nil {
			return elseCode.Call(c)
		}
		return Int(0), nil
	})
}

func MakeForEach(variable []byte, lst, code Callable) Callable {
	vn := string(variable)
	return callableFunc(func(c *Context) (Value, error) {
		list, err := callList(lst, c)
		if err != nil {
			return list, err
		}
		var res Value = Int(0)
		for _, val := range list.content {
			c.variables[vn] = val
			res, err = code.Call(c)
			if err != nil {
				return res, err
			}
		}
		return res, nil
	})
}
