package piscine

import "github.com/01-edu/z01"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 4; i++ {
		printStr("Player ")
		printInt(i + 1)
		printStr(": ")
		for j := 0; j < 3; j++ {
			printInt(deck[i*3+j])
			if j < 2 {
				printStr(", ")
			}
		}
		z01.PrintRune('\n')
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := []rune{}
	for n > 0 {
		digits = append([]rune{rune(n%10 + '0')}, digits...)
		n /= 10
	}
	for _, d := range digits {
		z01.PrintRune(d)
	}
}
