package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"time"
)

var timeOut = flag.Duration("timeout", DefaultTimeout, "defines the maximum runtime")
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

func loadCodeFromFile(fileName string) (*ParsedCode, error) {
	fCode, err := os.Open(fileName)
	if err != nil {
		return nil, err
	} else {
		defer fCode.Close()
		return Parse(fCode)
	}
}

func main() {
	flag.Usage = usageEx

	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		return
	}
	code, err := loadCodeFromFile(flag.Arg(0))

	if err != nil {
		flag.Usage()

		fmt.Println()
		fmt.Println(err)
		return
	}

	if *game {
		RunGame(code)
	} else {
		ctx := NewContext(*timeOut)

		if *input != "" {
			inpVal, err := ExecuteString(*input, 100*time.Millisecond)
			if err != nil {
				fmt.Fprintln(os.Stderr, "input value could not be parsed:", err)
				return
			}
			ctx.SetInput(inpVal)
		}
		res, err := code.Run(ctx)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Fprintln(os.Stdout, ToString(res))
	}
}
