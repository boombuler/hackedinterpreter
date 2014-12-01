# HackedInterpreter

HackedInterpreter is a featurecomplete interpreter for the H language which is used in http://hackedapp.com. This interpreter has supports C-style comments in code files, which the original language doesn't at the moment.

## Compiling
1. install [golang](http://golang.org/)
2. get the sources
3. run `go get code.google.com/p/gocc` to get the parser engine
4. run `go get github.com/nsf/termbox-go` to get the termbox package
5. run `go build`

### Modifying the Syntax
The syntax is declared in the `hacked.bnf` file and will be compiled by gocc.
To update the sourcecode run `gocc -p=".." hacked.bnf` and rebuild with `go build`

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

## Debugging
To start debugging start the program with `-debug` parameter. Once the debugger is attached it will wait before executing any code. On the right side you see the current state of each variable.

### Controls
You can move the code around with the arrow keys.
Once the debugger stopped code execution you can step over each single command by pressing `s`. If you want to continue the code execution till the next breakpoint press `c` to continue execution.
You can enter debugger commands by pressing `Enter` and enter one of the following commands:
* `bpx [Line]` will toggle a code breakpoint on the given line.
* `bpr [var name]` will toggle a memory breakpoint. Which will break as soon as any code has read the content of the given variable. Example: `bpr var_a` will break on code like `var_c = var_a * 3`
* `bpw [var name]` will toggle a memory breakpoint. Which will break as soon as any code has set the value of the given variable. Example: `bpw var_a` will break on code like `var_a = abs(var_b)`
* `eval [code]` will execute the given code. For example you can call `eval var_a = abs(var_b)` to set the current value of var_a. If the given code will is faulty nothing will happen, and you will not get any result of the given code.

The `Esc` key stops execution will exit the debugger