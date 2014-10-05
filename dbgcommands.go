package main

import (
	"./runtime"
	"./token"
	"regexp"
	"strconv"
)

func iterateExecutableTokens(c *runtime.Callable) <-chan *token.Token {
	result := make(chan *token.Token)
	go func() {
		var itChild func(cc *runtime.Callable)
		itChild = func(cc *runtime.Callable) {
			if cc.Children != nil {
				for _, ccc := range cc.Children {
					if ccc != nil {
						if ccc.Token != nil {
							result <- ccc.Token
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
	cmd_CodeBp = regexp.MustCompile("(?:bpx\\W+)(\\d+)")
)

func (ws *dbgWorkspace) handleCommand(cmd string) {
	if cmd_CodeBp.Match([]byte(cmd)) {
		lineStr := string(cmd_CodeBp.FindSubmatch([]byte(cmd))[1])
		line, _ := strconv.Atoi(lineStr)
		token := ws.findTokenByLine(line)
		if token != nil {
			ws.dbg.ToggleCodeBreakPoint(token)
		}
	}
}
