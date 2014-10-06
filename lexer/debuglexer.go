package lexer

import (
	"../token"
	"io/ioutil"
	"unicode/utf8"
)

// A copy of the Lexer which tries to fix a bug with token positions and keeps track of all tokens
// so they can be compared with the callables from the parser
type DebugLexer struct {
	src    []byte
	pos    int
	line   int
	column int
	Tokens []*token.Token
}

func NewDebugLexer(src []byte) *DebugLexer {
	lexer := &DebugLexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
		Tokens: make([]*token.Token, 0),
	}
	return lexer
}

func NewDebugLexerFile(fpath string) (*DebugLexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewDebugLexer(src), nil
}

func (this *DebugLexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if this.pos >= len(this.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = this.pos, this.line, this.column
		this.Tokens = append(this.Tokens, tok)
		return
	}
	start, end := this.pos, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if this.pos >= len(this.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(this.src[this.pos:])
			this.pos += size
		}

		// Production start
		if rune1 != -1 {
			state = TransTab[state](rune1)
		} else {
			state = -1
		}
		// Production end
		if state != -1 {
			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = this.pos
			case ActTab[state].Ignore != "":
				start = this.pos
				state = 0
				if start >= len(this.src) {
					tok.Type = token.EOF
				}
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
	this.Tokens = append(this.Tokens, tok)
	return
}
