package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			z01.PrintRune(rune('0' + a/10))
			z01.PrintRune(rune('0' + a%10))
			z01.PrintRune(' ')
			z01.PrintRune(rune('0' + b/10))
			z01.PrintRune(rune('0' + b%10))
			if !(a == 1 && b == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
