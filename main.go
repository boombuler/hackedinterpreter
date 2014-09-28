package main

import (
	"./lexer"
	"./parser"
	"./runtime"
	"fmt"
)

func main() {
	lex := lexer.NewLexer([]byte("f1(2, 3)"))
	p := parser.NewParser()
	res, err := p.Parse(lex)
	if err != nil {
		fmt.Println(err)
	} else {
		callable, ok := res.(runtime.Callable)
		if !ok {
			fmt.Println("result is not callable!")
		} else {
			ctx := runtime.NewContext()
			resVal, err := callable.Call(ctx)
			if err != nil {
				fmt.Println("Runime error: ", err)
			} else {
				fmt.Println("result:", resVal)
			}
		}
	}
}
