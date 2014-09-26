# HackedInterpreter

HackedInterpreter is a featurecomplete interpreter for the H language which is used in http://hackedapp.com

## Compiling
1. install [golang](http://golang.org/)
2. get the sources
3. run `go get github.com/boombuler/gold` to get the parser engine
4. run `go build`

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
The syntax is declared in the `hacked.grm` file and will be compiled to the `hacked.egt` file with [Gold Parser](http://goldparser.org/) after the EGT file (Enhanced Grammar Table) is created run the `packegt.go` from the utils directory. That will create the `egt.go` file which includes the binary data which can be used by the parser engine.