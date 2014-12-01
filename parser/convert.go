package parser

import (
	"../runtime"

	"../token"
)

func c(at Attrib) runtime.Callable {
	rt, ok := at.(runtime.Callable)
	if !ok {
		return nil
	}
	return rt
}

func p(at Attrib) *token.Token {
	tkn, ok := at.(*token.Token)
	if !ok {
		return nil
	}
	return tkn
}

func str(at Attrib) string {
	tkn, ok := at.(*token.Token)
	if !ok {
		return ""
	}
	return string(tkn.Lit)
}

func assignable(at Attrib) runtime.Assignable {
	ass, ok := at.(runtime.Assignable)
	if !ok {
		return nil
	}
	return ass
}
