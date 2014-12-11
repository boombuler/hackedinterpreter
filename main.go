package main

//go:generate gocc -p=github.com/boombuler/hackedinterpreter hacked.bnf

import (
	"flag"
	"fmt"
	"github.com/boombuler/hackedinterpreter/runtime"
	"github.com/boombuler/hackedinterpreter/server"
	"os"
	"path"
)

var timeOut = flag.Duration("timeout", runtime.DefaultTimeout, "defines the maximum runtime")
var game = flag.Bool("game", false, "run the code in freestyle mode. \"timeout\" flag will be ignored!")
var debug = flag.Bool("debug", false, "allow the code to be debugged.")

func usageEx() {
	prName := path.Base(os.Args[0])
	fmt.Println("Usage of", prName, ":")
	fmt.Println("   ", prName, "[flags]", "code.file")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usageEx

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		return
	}

	code, lex, err := fromFile(flag.Arg(0))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var srv *server.Server
	if *game {
		srv, err = server.NewGameServer(code, lex, *debug)
	} else {
		srv, err = server.NewServer(code, lex, *timeOut, *debug)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(srv.Addr)
	srv.Wait()
}
