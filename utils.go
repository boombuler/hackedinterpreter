package main

import (
	"./lexer"
	"./parser"
	"./runtime"
	"time"
)

func fromString(code string) (*runtime.Callable, error) {
	lex := lexer.NewLexer([]byte(code))
	p := parser.NewParser()
	res, err := p.Parse(lex)

	if err != nil {
		return nil, err
	} else {
		return res.(*runtime.Callable), nil
	}
}

func fromFile(fileName string) (*runtime.Callable, error) {
	lex, err := lexer.NewLexerFile(fileName)
	if err != nil {
		return nil, err
	}
	p := parser.NewParser()
	res, err := p.Parse(lex)
	if err != nil {
		return nil, err
	} else {
		return res.(*runtime.Callable), nil
	}
}

func execString(code string, timeOut time.Duration) (runtime.Value, error) {
	call, err := fromString(code)
	if err != nil {
		return nil, err
	}
	return exec(call, timeOut)
}

func exec(c *runtime.Callable, timeOut time.Duration) (runtime.Value, error) {
	ctx := runtime.NewContext(timeOut)
	return c.Call(ctx)
}
