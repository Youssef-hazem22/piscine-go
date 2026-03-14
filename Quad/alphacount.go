package piscine

func AlphaCount(s string) int {
	counter := 0
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			counter++
		}
	}
	return counter
}
