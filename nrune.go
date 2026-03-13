package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	letter := 1
	for _, word := range s {
		if letter == n {
			return word
		}
		letter++
	}

	return 0
}
