package main

import "github.com/01-edu/z01"

func main() {
	four := 4 + '0'
	two := 2 + '0'
	one := 1 + '0'

	digits := []rune{'x', ' ', '=', ' ', four, two, ',', ' ', 'y', ' ', '=', ' ', two, one, '\n'}
	for _, d := range digits {
		z01.PrintRune(d)
	}
}
