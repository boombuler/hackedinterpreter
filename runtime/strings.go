package runtime

import (
	"../token"
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func NewConstStr(text string, p *token.Token) (Callable, error) {
	s, err := strconv.Unquote(text)
	if err != nil {
		return nil, err
	}
	return newCallable(p, func(c *Context) (Value, error) {
		return s, nil
	}), nil
}

func strLength(s string) int {
	return utf8.RuneCountInString(s)
}

func strIndexInRange(s string, i int) bool {
	return i >= 0 && i < strLength(s)
}

func strGetRuneByIdx(str string, idx int) string {
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		if idx == 0 {
			return fmt.Sprintf("%c", r)
		}
		idx--
		str = str[size:]
	}
	return ""
}

func strLeft(str string, l int) string {
	b := new(bytes.Buffer)
	for len(str) > 0 && l > 0 {
		r, size := utf8.DecodeRuneInString(str)
		b.WriteRune(r)
		l--
		str = str[size:]
	}
	return b.String()
}

func strRight(str string, l int) string {
	skip := strLength(str) - l
	for len(str) > 0 && skip > 0 {
		_, size := utf8.DecodeRuneInString(str)
		skip--
		str = str[size:]
	}
	return str
}
