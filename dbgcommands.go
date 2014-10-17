package main

import (
	"./runtime"
	"./token"
	"regexp"
	"strconv"
)

func iterateExecutableTokens(c runtime.Callable) <-chan *token.Token {
	result := make(chan *token.Token)
	go func() {
		var itChild func(cc runtime.Callable)
		itChild = func(cc runtime.Callable) {
			meta := cc.Meta()
			if meta != nil && meta.Children != nil {
				for _, ccc := range meta.Children {
					if ccc != nil {
						ccm := ccc.Meta()
						if ccm != nil && ccm.Token != nil {
							result <- ccm.Token
						}
						itChild(ccc)
					}
				}
			}
		}
		itChild(c)
		close(result)
	}()
	return result
}

func (ws *dbgWorkspace) findTokenByLine(line int) (t *token.Token) {
	for tkn := range iterateExecutableTokens(ws.code) {
		if tkn.Line == line && t == nil {
			t = tkn
		}
	}
	return
}

var (
	cmdCodeBp = regexp.MustCompile("(?:bpx\\W+)(\\d+)")
	cmdEval   = regexp.MustCompile("(?:eval\\W+)(.*)")
)

func (ws *dbgWorkspace) handleCommand(cmd string) {
	bcmd := []byte(cmd)
	if cmdCodeBp.Match(bcmd) {
		lineStr := string(cmdCodeBp.FindSubmatch(bcmd)[1])
		line, _ := strconv.Atoi(lineStr)
		token := ws.findTokenByLine(line)
		if token != nil {
			ws.dbg.ToggleCodeBreakPoint(token)
		}
		return
	}
	if cmdEval.Match(bcmd) {
		code := string(cmdEval.FindSubmatch(bcmd)[1])
		executable, err := fromString(code)
		if err == nil {
			ws.dbg.Eval(executable)
		}
		return
	}
}
