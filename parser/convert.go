package parser

import (
	"../runtime"
	"../token"
)

func c(at Attrib) runtime.Callable {
	rt := at.(runtime.Callable)
	return rt
}

func str(at Attrib) string {
	tkn, ok := at.(*token.Token)
	if !ok {
		return ""
	}
	return string(tkn.Lit)
}
