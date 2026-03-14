package piscine

func LastRune(s string) rune {
	var last rune
	for _, letter := range s {
		last = letter
	}
	return last
}
