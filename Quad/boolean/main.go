package main

import (
	"os"

	"github.com/01-edu/z01"
)

const (
	evenmsg = "I have an even number of arguments"
	oddmsg  = "I have an odd number of arguments"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func sure(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}

func isEven(nbr int) bool {
	if sure(nbr) == 1 {
		return true
	}
	return false
}

func main() {
	arg := len(os.Args[1:])
	if isEven(arg) {
		printStr(evenmsg)
	} else {
		printStr(oddmsg)
	}
}
