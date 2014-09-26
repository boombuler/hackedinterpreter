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
	codeFile := flag.Arg(0)
	fCode, err := os.Open(codeFile)
	if err != nil {
		flag.Usage()

		fmt.Println()
		fmt.Println(err)
	}
	defer fCode.Close()

	ctx := NewContext(*timeOut)

	if *input != "" {
		inpVal, err := ExecuteString(*input, 100*time.Millisecond)
		if err != nil {
			fmt.Fprintln(os.Stderr, "input value could not be parsed:", err)
			return
		}
		ctx.SetInput(inpVal)
	}
	res, err := ExecuteReader(fCode, ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Fprintln(os.Stdout, ToString(res))
}
