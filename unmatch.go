package piscine

func Unmatch(a []int) int {
	for _, ch := range a {
		count := 0
		for _, yh := range a {
			if ch == yh {
				count++
			}
		}
		if count%2 != 0 {
			return ch
		}
	}
	return -1
}
