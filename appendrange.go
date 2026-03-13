package piscine

func AppendRange(min, max int) []int {
	var s []int
	if min >= max {
		s = nil
		return s
	}
	for i := min; i < max; i++ {
		s = append(s, i)
	}
	return s
}
