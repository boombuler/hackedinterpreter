
package token

import(
	"fmt"
)

type Token struct {
	Type
	Lit []byte
	Pos
}

type Type int

const(
	INVALID Type = iota
	EOF
)

type Pos struct {
	Offset int
	Line int
	Column int
}

func (this Pos) String() string {
	return fmt.Sprintf("Pos(offset=%d, line=%d, column=%d)", this.Offset, this.Line, this.Column)
}

type TokenMap struct {
	typeMap  []string
	idMap map[string]Type
}

func (this TokenMap) Id(tok Type) string {
	if int(tok) < len(this.typeMap) {
		return this.typeMap[tok]
	}
	return "unknown"
}

func (this TokenMap) Type(tok string) Type {
	if typ, exist := this.idMap[tok]; exist {
		return typ
	}
	return INVALID
}

func (this TokenMap) TokenString(tok *Token) string {
	//TODO: refactor to print pos & token string properly
	return fmt.Sprintf("%s(%d,%s)", this.Id(tok.Type), tok.Type, tok.Lit)
}

func (this TokenMap) StringType(typ Type) string {
	return fmt.Sprintf("%s(%d)", this.Id(typ), typ)
}

var TokMap = TokenMap{
	typeMap: []string{
		"INVALID",
		"$",
		"error",
		"var",
		"input",
		"true",
		"false",
		"(",
		")",
		"cust_fn_name",
		"int",
		"*",
		"/",
		"+",
		"-",
		">",
		"<",
		"==",
		"!=",
		"&&",
		"||",
		"[",
		"]",
		"=",
		",",
		"fn_name",
		".",
		"{",
		"}",
		"function",
		":",
		"return",
		";",
		"if",
		"else",
		"while",
		"foreach",
		"in",
		"->",
	},

	idMap: map[string]Type {
		"INVALID": 0,
		"$": 1,
		"error": 2,
		"var": 3,
		"input": 4,
		"true": 5,
		"false": 6,
		"(": 7,
		")": 8,
		"cust_fn_name": 9,
		"int": 10,
		"*": 11,
		"/": 12,
		"+": 13,
		"-": 14,
		">": 15,
		"<": 16,
		"==": 17,
		"!=": 18,
		"&&": 19,
		"||": 20,
		"[": 21,
		"]": 22,
		"=": 23,
		",": 24,
		"fn_name": 25,
		".": 26,
		"{": 27,
		"}": 28,
		"function": 29,
		":": 30,
		"return": 31,
		";": 32,
		"if": 33,
		"else": 34,
		"while": 35,
		"foreach": 36,
		"in": 37,
		"->": 38,
	},
}

