package main

import (
	"github.com/boombuler/gold"
	"io"
)

type ParsedCode gold.Token

func Parse(reader io.Reader) (*ParsedCode, error) {
	res, err := parser.Parse(reader, true)
	if err != nil {
		return nil, err
	}
	return (*ParsedCode)(res), nil
}

func (pc *ParsedCode) Run(c *Context) (Value, error) {
	result := make(chan Value, 1)
	resErr := make(chan error, 1)

	go func() {
		r, e := Exec((*gold.Token)(pc), c)
		resErr <- e
		result <- r
		close(resErr)
		close(result)
	}()

	return <-result, <-resErr
}
