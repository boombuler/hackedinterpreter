## Introduction
To start debugging start the program with `-debug` parameter. Once the debugger is attached it will wait before executing any code. On the right side you see the current state of each variable.

## Controls
You can move the code around with the arrow keys.
Once the debugger stopped code execution you can step over each single command by pressing `s`. If you want to continue the code execution till the next breakpoint press `c` to continue execution.
You can enter debugger commands by pressing `Enter` and enter one of the following commands:
* `bpx [Line]` will toggle a code breakpoint on the given line.
* `bpr [var name]` will toggle a memory breakpoint. Which will break as soon as any code reads the content of the given variable. Example: `bpr var_a` will break on code like `var_c = var_a * 3`
* `bpw [var name]` will toggle a memory breakpoint. Which will break as soon as any code set the value of the given variable. Example: `bpw var_a` will break on code like `var_a = abs(var_b)`
* `eval [code]` will execute the given code. For example you can call `eval var_a = abs(var_b)` to set the current value of var_a. If the given code will is faulty nothing will happen, and you will not get any result of the given code.

The `Esc` key stops execution will exit the debugger
