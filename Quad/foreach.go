package piscine

func PrintNbr(n int) {
}

func ForEach(f func(int), a []int) {
	for _, v := range a {
		f(v)
	}
}
