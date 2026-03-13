package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	youssef := true
	hazem := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			youssef = false
		}
		if f(a[i], a[i+1]) < 0 {
			hazem = false
		}
	}
	return youssef || hazem
}
