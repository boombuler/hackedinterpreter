package server

import (
	"bytes"
	"errors"
	"fmt"
	pe "github.com/boombuler/hackedinterpreter/errors"
	"github.com/boombuler/hackedinterpreter/lexer"
	"github.com/boombuler/hackedinterpreter/parser"
	"github.com/boombuler/hackedinterpreter/runtime"
	"io/ioutil"
	"net/http"
)

type codeLoader struct {
	*Server
}

func (cl *codeLoader) RegisterHandlers(s *Server, m *http.ServeMux) {
	cl.Server = s
	m.HandleFunc("/loadcode", cl.handleLoadCode)
}

func (cl *codeLoader) CodeLoaded() {
	// nothing to do.
}

func checkParseError(res interface{}, err error) (runtime.Callable, error) {
	if err != nil {
		return nil, err
	}
	if call, isCallable := res.(runtime.Callable); isCallable {
		return call, nil
	}

	if E, isError := res.(*pe.Error); isError {
		w := new(bytes.Buffer)
		fmt.Fprintf(w, "Error")
		if E.Err != nil {
			fmt.Fprintf(w, " %s\n", E.Err)
		} else {
			fmt.Fprintf(w, "\n")
		}
		fmt.Fprintf(w, "Invalid token: \"%s\"", E.ErrorToken.Lit)
		fmt.Fprintf(w, " at Line %d / Column %d", E.ErrorToken.Pos.Line, E.ErrorToken.Pos.Column)

		return nil, errors.New(w.String())
	}
	panic(fmt.Sprintf("Unknown parse result: %v", res))
}

func (cl *codeLoader) handleLoadCode(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("codefile")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	lex := lexer.NewDebugLexer(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	p := parser.NewParser()
	code, err := checkParseError(p.Parse(lex))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	cl.code = code
	cl.tokens = lex.Tokens
	cl.ctx = runtime.NewContext(runtime.DefaultTimeout)
	for _, mod := range cl.modules {
		mod.CodeLoaded()
	}
	http.Redirect(w, r, "/index", http.StatusFound)
}
