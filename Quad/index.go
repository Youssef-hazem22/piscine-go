package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	if len(s) < 1 || len(toFind) < 1 {
		return -1
	}
	if len(toFind) > len(s) || len(toFind) == len(s) {
		return -1
	}

	counter := 0
	srune := []rune(s)
	frune := []rune(toFind)[0] // now it's safe
	match := false

	for _, ch := range srune {
		if ch == frune {
			match = true
			break
		}
		counter++
	}

	if match {
		return counter
	}
	return -1
}
