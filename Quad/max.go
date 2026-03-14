package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	num := a[0]
	for _, ch := range a {
		if ch > num {
			num = ch
		}
	}
	return num
}
