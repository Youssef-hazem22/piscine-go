package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programPath := os.Args[0]
	name := ""

	for i := len(programPath) - 1; i >= 0; i-- {
		if programPath[i] == '/' {
			name = programPath[i+1:]
			break
		}
	}

	if name == "" {
		name = programPath
	}

	for _, r := range name {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
