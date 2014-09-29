# HackedInterpreter

HackedInterpreter is a featurecomplete interpreter for the H language which is used in http://hackedapp.com. This interpreter has supports C-style comments in code files, which the original language doesn't at the moment. 

## Compiling
1. install [golang](http://golang.org/)
2. get the sources
3. run `go get code.google.com/p/gocc` to get the parser engine
4. run `go get github.com/nsf/termbox-go` to get the termbox package
5. run `go build`

## Executing Code
Simply call the compiled binary with the code you want to execute.  By default all code execution will be terminated, if the program runs more than 5 seconds.

### Examples
* Simple execution.
  `hackedinterpreter example.h`
* With timeout set.
  `hackedinterpreter -timeout=10s example.h`
* With given input.
  the input is parsed as H-Syntax too. But the parsing will be canceled if it
  doesn't parse in 100 milliseconds.
  `hackedinterpreter -input="[1, 2, 3]" example.h`
* Start a game. Controls: Keyboard Up / Down / Left / Right / A / B
  `hackedinterpreter -game game.h`

## Modifying the Syntax
The syntax is declared in the `hacked.bnf` file and will be compiled by gocc.
To update the sourcecode run `gocc -p=".." hacked.bnf` and rebuild with `go build`