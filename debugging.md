## Introduction
To start debugging start the program with `-debug` parameter. Once the debugger is attached it will wait before executing any code. On the right side you see the current state of each variable.

## Controls
You can move the code around with the arrow keys.
Once the debugger stopped code execution you can step over each single command by pressing `s`. If you want to continue the code execution till the next breakpoint press `c` to continue execution.
You can enter debugger commands by pressing `Enter` and enter one of the following commands:
* `bpx [Line]` will toggle a code breakpoint on the given line.