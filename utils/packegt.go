package main

import (
	"io"
	"os"
)

const hextable = "0123456789abcdef"
const bufferSize = 1024

func main() {
	egt, _ := os.Open("..\\hacked.egt")
	defer egt.Close()
	gof, _ := os.Create("..\\egt.go")
	defer gof.Close()
	gof.WriteString("package main\n var EGTFile =\"")

	buffer := make([]byte, bufferSize, bufferSize)

	for read, err := egt.Read(buffer); err != io.EOF; read, err = egt.Read(buffer) {
		for i := 0; i < read; i++ {
			gof.WriteString("\\x")
			v := buffer[i]
			gof.Write([]byte{hextable[v>>4], hextable[v&0x0f]})
		}
	}

	gof.WriteString("\"")
}
