
package lexer

import (
	
	// "fmt"
	// "../util"
	
	"io/ioutil"
	"unicode/utf8"
	"../token"
)

const(
	NoState = -1
	NumStates = 85
	NumSymbols = 102
) 

type Lexer struct {
	src             []byte
	pos             int
	line            int
	column          int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (this *Lexer) Scan() (tok *token.Token) {
	
	// fmt.Printf("Lexer.Scan() pos=%d\n", this.pos)
	
	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
	
		// fmt.Printf("\tpos=%d, line=%d, col=%d, state=%d\n", this.pos, this.line, this.column, state)
	
		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}
		switch rune1 {
		case '\n':
			this.line++
			this.column = 1
		case '\r':
			this.column = 1
		case '\t':
			this.column += 4
		default:
			this.column++
		}

	
		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end

		// Debug start
		// nextState := -1
		// if rune1 != -1 {
		// 	nextState = TransTab[state](rune1)
		// }
		// fmt.Printf("\tS%d, : tok=%s, rune == %s(%x), next state == %d\n", state, token.TokMap.Id(tok.Type), util.RuneToString(rune1), rune1, nextState)
		// fmt.Printf("\t\tpos=%d, size=%d, start=%d, end=%d\n", this.pos, size, start, end)
		// if nextState != -1 {
		// 	fmt.Printf("\t\taction:%s\n", ActTab[nextState].String())
		// }
		// state = nextState
		// Debug end
	

		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				// fmt.Printf("\t Accept(%s), %s(%d)\n", string(act), token.TokMap.Id(tok), tok)
				end = this.pos
			case ActTab[state].Ignore != "":
				// fmt.Printf("\t Ignore(%s)\n", string(act))
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = this.pos
			}
		}
	}
	if end > start {
		this.pos = end
		tok.Lit = this.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset = start
	tok.Pos.Column = this.column
	tok.Pos.Line = this.line
	return
}

func (this *Lexer) Reset() {
	this.pos = 0
}

/*
Lexer symbols:
0: 'v'
1: 'a'
2: 'r'
3: '_'
4: 'f'
5: 'i'
6: 'n'
7: 'p'
8: 'u'
9: 't'
10: 't'
11: 'r'
12: 'u'
13: 'e'
14: 'f'
15: 'a'
16: 'l'
17: 's'
18: 'e'
19: '('
20: ')'
21: '-'
22: '!'
23: '*'
24: '/'
25: '+'
26: '>'
27: '<'
28: '='
29: '='
30: '!'
31: '='
32: '&'
33: '&'
34: '|'
35: '|'
36: '['
37: ']'
38: '='
39: ','
40: '.'
41: '{'
42: '}'
43: 'f'
44: 'u'
45: 'n'
46: 'c'
47: 't'
48: 'i'
49: 'o'
50: 'n'
51: ':'
52: 'r'
53: 'e'
54: 't'
55: 'u'
56: 'r'
57: 'n'
58: ';'
59: 'i'
60: 'f'
61: 'e'
62: 'l'
63: 's'
64: 'e'
65: 'w'
66: 'h'
67: 'i'
68: 'l'
69: 'e'
70: 'f'
71: 'o'
72: 'r'
73: 'e'
74: 'a'
75: 'c'
76: 'h'
77: 'i'
78: 'n'
79: '-'
80: '>'
81: '_'
82: '/'
83: '/'
84: '\n'
85: '\r'
86: '/'
87: '*'
88: '*'
89: '*'
90: '/'
91: ' '
92: '\t'
93: '\n'
94: '\r'
95: '0'-'9'
96: 'a'-'o'
97: 'a'-'z'
98: '1'-'9'
99: '0'-'6'
100: 'a'-'z'
101: .

*/
