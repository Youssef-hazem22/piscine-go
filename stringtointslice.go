package piscine

func StringToIntSlice(str string) []int {
	var slice []int
	for _, ch := range str {
		slice = append(slice, int(ch))
	}
	return slice
}
