
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
	NumStates = 91
	NumSymbols = 109
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
5: '"'
6: '"'
7: 'i'
8: 'n'
9: 'p'
10: 'u'
11: 't'
12: 't'
13: 'r'
14: 'u'
15: 'e'
16: 'f'
17: 'a'
18: 'l'
19: 's'
20: 'e'
21: '('
22: ')'
23: '['
24: ']'
25: '+'
26: '+'
27: '-'
28: '-'
29: '-'
30: '!'
31: '*'
32: '/'
33: '+'
34: '>'
35: '<'
36: '='
37: '='
38: '!'
39: '='
40: '&'
41: '&'
42: '|'
43: '|'
44: '='
45: ','
46: '.'
47: '{'
48: '}'
49: 'f'
50: 'u'
51: 'n'
52: 'c'
53: 't'
54: 'i'
55: 'o'
56: 'n'
57: ':'
58: 'r'
59: 'e'
60: 't'
61: 'u'
62: 'r'
63: 'n'
64: ';'
65: 'i'
66: 'f'
67: 'e'
68: 'l'
69: 's'
70: 'e'
71: 'w'
72: 'h'
73: 'i'
74: 'l'
75: 'e'
76: 'f'
77: 'o'
78: 'r'
79: 'e'
80: 'a'
81: 'c'
82: 'h'
83: 'i'
84: 'n'
85: '-'
86: '>'
87: '_'
88: '\'
89: '/'
90: '/'
91: '\n'
92: '\r'
93: '/'
94: '*'
95: '*'
96: '*'
97: '/'
98: ' '
99: '\t'
100: '\n'
101: '\r'
102: '0'-'9'
103: 'a'-'o'
104: 'a'-'z'
105: '1'-'9'
106: '0'-'6'
107: 'a'-'z'
108: .

*/
