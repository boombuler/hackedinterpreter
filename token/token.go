
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
		"var",
		"input",
		"true",
		"false",
		"(",
		")",
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
		"[]",
		"fn_name",
		"cust_fn_name",
		"()",
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
		"var": 2,
		"input": 3,
		"true": 4,
		"false": 5,
		"(": 6,
		")": 7,
		"int": 8,
		"*": 9,
		"/": 10,
		"+": 11,
		"-": 12,
		">": 13,
		"<": 14,
		"==": 15,
		"!=": 16,
		"&&": 17,
		"||": 18,
		"[": 19,
		"]": 20,
		"=": 21,
		",": 22,
		"[]": 23,
		"fn_name": 24,
		"cust_fn_name": 25,
		"()": 26,
		".": 27,
		"{": 28,
		"}": 29,
		"function": 30,
		":": 31,
		"return": 32,
		";": 33,
		"if": 34,
		"else": 35,
		"while": 36,
		"foreach": 37,
		"in": 38,
		"->": 39,
	},
}

