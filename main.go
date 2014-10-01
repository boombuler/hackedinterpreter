package main

import (
	"./runtime"
	"flag"
	"fmt"
	"os"
	"path"
	"time"
)

var timeOut = flag.Duration("timeout", runtime.DefaultTimeout, "defines the maximum runtime")
var input = flag.String("input", "", "defines what should be used as Input value.")
var game = flag.Bool("game", false, "run the code in freestyle mode. Input and timeout will be ignored!")

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
	code, err := fromFile(flag.Arg(0))

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if *game {
		RunGame(code)
	} else {
		ctx := runtime.NewContext(*timeOut)

		if *input != "" {
			inpCode, err := fromString(*input)
			if err != nil {
				fmt.Fprintln(os.Stderr, "input value could not be parsed:", err)
				return
			}
			inpVal, err := exec(inpCode, 100*time.Millisecond)
			if err != nil {
				fmt.Fprintln(os.Stderr, "input value could not be parsed:", err)
				return
			}
			ctx.SetInput(inpVal)
		}
		res, err := code.Call(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Fprintln(os.Stdout, runtime.ToString(res))
		}
	}
}
