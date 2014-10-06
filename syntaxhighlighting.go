package main

import (
	"./token"
	"github.com/nsf/termbox-go"
)

const hlKeyword = termbox.AttrBold | termbox.ColorBlue
const hlVariable = termbox.ColorCyan
const hlConst = termbox.ColorMagenta

var highlightMap map[string]termbox.Attribute = map[string]termbox.Attribute{
	"function":     hlKeyword,
	"return":       hlKeyword,
	"if":           hlKeyword,
	"else":         hlKeyword,
	"while":        hlKeyword,
	"foreach":      hlKeyword,
	"in":           hlKeyword,
	"->":           hlKeyword,
	":":            hlKeyword,
	"var":          hlVariable,
	"input":        hlVariable,
	"cust_fn_name": hlVariable,
	"fn_name":      hlVariable,
	"true":         hlConst,
	"false":        hlConst,
	"int":          hlConst,
}

func getForeground(t *token.Token) termbox.Attribute {
	tokId := token.TokMap.Id(t.Type)
	att, ok := highlightMap[tokId]
	if ok {
		return att
	}
	return termbox.ColorWhite
}
