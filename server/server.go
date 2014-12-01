package server

import (
	"../lexer"
	"../runtime"
	"../token"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Server struct {
	code   runtime.Callable
	ctx    *runtime.Context
	tokens []*token.Token

	Addr       string
	stopServer func() error
	closeChan  chan struct{}
}

type TokenExport struct {
	token.Pos
	Text string `json:",omitempty"`
	Type string
}

func NewServer(code runtime.Callable, lex *lexer.DebugLexer, timeout time.Duration, debuggable bool) (*Server, error) {
	return newServer(code, lex, debuggable, func(s *Server, m *http.ServeMux) {

	})
}

func newServer(code runtime.Callable, lex *lexer.DebugLexer, debuggable bool, addHandlers func(s *Server, m *http.ServeMux)) (*Server, error) {
	server := &Server{
		code:      code,
		ctx:       runtime.NewContext(runtime.NoTimeout),
		tokens:    lex.Tokens,
		closeChan: make(chan struct{}),
	}

	go func() {
		<-server.closeChan
		if server.stopServer != nil {
			if err := server.stopServer(); err != nil {
				panic(err.Error())
			}
		}
	}()

	var addr net.Addr
	addr, server.stopServer = serveHttp(func(m *http.ServeMux) {
		if debuggable {
			m.HandleFunc("/code", server.serveCode)
		}

		m.HandleFunc("/stop", func(resp http.ResponseWriter, req *http.Request) {
			close(server.closeChan)
		})
		addHandlers(server, m)
	})

	server.Addr = fmt.Sprintf("http://localhost:%v/gamewnd.html", (addr.(*net.TCPAddr)).Port)

	return server, nil
}

func (s *Server) Wait() {
	<-s.closeChan
}

func (s *Server) serveCode(resp http.ResponseWriter, req *http.Request) {
	tex := make([]*TokenExport, 0, len(s.tokens))
	for _, t := range s.tokens {
		if len(t.Lit) == 0 {
			continue
		}
		x := &TokenExport{
			t.Pos,
			string(t.Lit),
			token.TokMap.Id(t.Type),
		}

		if x.Text == x.Type {
			x.Text = ""
		}
		tex = append(tex, x)
	}
	err := json.NewEncoder(resp).Encode(tex)
	if err != nil {
		panic(err)
	}
}

func serveHttp(regHandlers func(*http.ServeMux)) (net.Addr, func() error) {
	l, err := net.ListenTCP("tcp", nil)
	if err != nil {
		panic(err.Error())
	}
	srvMux := http.NewServeMux()
	srv := &http.Server{
		Handler: srvMux,
	}
	regHandlers(srvMux)

	go func() {
		if err = srv.Serve(l); err != nil {
			panic(err.Error())
		}
	}()

	return l.Addr(), l.Close
}
