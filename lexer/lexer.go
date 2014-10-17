
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
	NumStates = 87
	NumSymbols = 106
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
21: '['
22: ']'
23: '+'
24: '+'
25: '-'
26: '-'
27: '-'
28: '!'
29: '*'
30: '/'
31: '+'
32: '>'
33: '<'
34: '='
35: '='
36: '!'
37: '='
38: '&'
39: '&'
40: '|'
41: '|'
42: '='
43: ','
44: '.'
45: '{'
46: '}'
47: 'f'
48: 'u'
49: 'n'
50: 'c'
51: 't'
52: 'i'
53: 'o'
54: 'n'
55: ':'
56: 'r'
57: 'e'
58: 't'
59: 'u'
60: 'r'
61: 'n'
62: ';'
63: 'i'
64: 'f'
65: 'e'
66: 'l'
67: 's'
68: 'e'
69: 'w'
70: 'h'
71: 'i'
72: 'l'
73: 'e'
74: 'f'
75: 'o'
76: 'r'
77: 'e'
78: 'a'
79: 'c'
80: 'h'
81: 'i'
82: 'n'
83: '-'
84: '>'
85: '_'
86: '/'
87: '/'
88: '\n'
89: '\r'
90: '/'
91: '*'
92: '*'
93: '*'
94: '/'
95: ' '
96: '\t'
97: '\n'
98: '\r'
99: '0'-'9'
100: 'a'-'o'
101: 'a'-'z'
102: '1'-'9'
103: '0'-'6'
104: 'a'-'z'
105: .

*/
