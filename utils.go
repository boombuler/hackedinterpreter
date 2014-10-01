package main

import (
	pe "./errors"
	"./lexer"
	"./parser"
	"./runtime"
	"bytes"
	"errors"
	"fmt"
	"time"
)

func CheckParseError(res interface{}, err error) (*runtime.Callable, error) {
	if err != nil {
		return nil, err
	}
	if call, isCallable := res.(*runtime.Callable); isCallable {
		return call, nil
	}

	if E, isError := res.(*pe.Error); isError {
		w := new(bytes.Buffer)
		fmt.Fprintf(w, "Error")
		if E.Err != nil {
			fmt.Fprintf(w, " %s\n", E.Err)
		} else {
			fmt.Fprintf(w, "\n")
		}
		fmt.Fprintf(w, "Invalid token: \"%s\"", E.ErrorToken.Lit)
		fmt.Fprintf(w, " at Line %d / Column %d", E.ErrorToken.Pos.Line, E.ErrorToken.Pos.Column)

		return nil, errors.New(w.String())
	}
	panic(fmt.Sprintf("Unknown parse result: %v", res))
}

func fromString(code string) (*runtime.Callable, error) {
	lex := lexer.NewLexer([]byte(code))
	p := parser.NewParser()
	res, err := CheckParseError(p.Parse(lex))

	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func fromFile(fileName string) (*runtime.Callable, error) {
	lex, err := lexer.NewLexerFile(fileName)
	if err != nil {
		return nil, err
	}
	p := parser.NewParser()
	res, err := CheckParseError(p.Parse(lex))
	if err != nil {
		return nil, err
	} else {
		return res, nil
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
