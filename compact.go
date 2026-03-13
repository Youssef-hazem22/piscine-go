package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	sl := []string{}

	for _, v := range s {
		if v != "" {
			sl = append(sl, v)
		}
	}

	*ptr = sl
	return len(sl)
}
