package main

//go:generate gocc -p=github.com/boombuler/hackedinterpreter hacked.bnf
//go:generate go generate github.com/boombuler/hackedinterpreter/server

import (
	"fmt"
	"github.com/boombuler/hackedinterpreter/server"
	"os"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	openBrowser(srv.Addr)
	srv.Wait()
}
