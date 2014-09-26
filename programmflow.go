package main

import (
	"github.com/boombuler/gold"
)

func ExecForEach(t *gold.Token, c *Context) (Value, error) {
	// <ForEachLoop> ::= foreach Variable in <Expression> <CodeBlock>
	varName, err := varNameFromToken(t.Tokens[1])
	if err != nil {
		return nil, err
	}
	lst, err := ExecWantList(t.Tokens[3], c)
	if err != nil {
		return nil, err
	}
	var res Value = nil
	for _, val := range lst.content {
		c.variables[varName] = val
		res, err = Exec(t.Tokens[4], c)

		if c.forceExit() {
			return c.result, c.err
		}

		if err != nil {
			return res, err
		}
	}
	return res, nil
}

func ExecIfBlock(t *gold.Token, c *Context) (Value, error) {
	// <IfBlock> ::= if <Expression> <CodeBlock> <ElseBlock>
	isTrue, err := ExecWantBool(t.Tokens[1], c)
	if err != nil {
		return isTrue, err
	}
	if isTrue {
		return Exec(t.Tokens[2], c)
	} else {
		tElse := t.Tokens[3]
		if tElse.Rule == Rule_ElseblockElse {
			return Exec(tElse.Tokens[1], c)
		}
	}
	return 0, nil
}

func ExecWhileLoop(t *gold.Token, c *Context) (Value, error) {
	// <WhileLoop> ::= while <Expression> <CodeBlock>
	var res Value = 0
	for {
		isTrue, err := ExecWantBool(t.Tokens[1], c)
		if err != nil {
			return 0, err
		}
		if isTrue {
			res, err = Exec(t.Tokens[2], c)
			if c.forceExit() {
				return c.result, c.err
			}

			if err != nil {
				return res, err
			}
		} else {
			break
		}
	}
	return res, nil
}

func ExecCodeBlock(t *gold.Token, c *Context) (Value, error) {
	if c.forceExit() {
		return c.result, c.err
	}
	return Exec(t.Tokens[1], c)
}

func ExecStatements(t *gold.Token, c *Context) (Value, error) {
	if t.Rule != Rule_Statements {
		return Exec(t.Tokens[0], c)
	} else {
		v, err := Exec(t.Tokens[0], c)
		if err != nil {
			return v, err
		}
		if c.forceExit() {
			return c.result, c.err
		}
		return Exec(t.Tokens[1], c)
	}
}
