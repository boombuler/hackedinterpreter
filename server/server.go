package server

//go:generate esc -o=content.go -pkg=server -prefix=content ./content/

import (
	"fmt"
	"github.com/boombuler/hackedinterpreter/runtime"
	"github.com/boombuler/hackedinterpreter/token"
	"net"
	"net/http"
)

type Server struct {
	code   runtime.Callable
	ctx    *runtime.Context
	tokens []*token.Token

	Addr       string
	stopServer func() error
	closeChan  chan struct{}

	modules []module
}

type module interface {
	RegisterHandlers(s *Server, m *http.ServeMux)
	CodeLoaded()
}

func NewServer() (*Server, error) {
	server := &Server{
		code:      nil,
		ctx:       nil,
		tokens:    nil,
		closeChan: make(chan struct{}),
		modules: []module{
			new(indexMod),
			new(codeServer),
			new(codeLoader),
		},
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
		m.HandleFunc("/stop", func(resp http.ResponseWriter, req *http.Request) {
			close(server.closeChan)
		})

		for _, mod := range server.modules {
			mod.RegisterHandlers(server, m)
		}

		m.Handle("/", http.FileServer(FS(false)))
	})

	server.Addr = fmt.Sprintf("http://localhost:%v/index", (addr.(*net.TCPAddr)).Port)

	return server, nil
}

func (s *Server) Wait() {
	<-s.closeChan
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
