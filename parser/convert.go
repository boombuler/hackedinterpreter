package parser

import (
	"../runtime"
)

func c(at Attrib) runtime.Callable {
	rt := at.(runtime.Callable)
	return rt
}
