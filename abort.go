package piscine

func Abort(a, b, c, d, e int) int {
	array := [5]int{a, b, c, d, e}
	for i := 0; i < len(array); i++ {
		for x := 0; x < len(array)-1; x++ {
			if array[x] > array[x+1] {
				array[x], array[x+1] = array[x+1], array[x]
			}
		}
	}
	return array[len(array)/2]
}
