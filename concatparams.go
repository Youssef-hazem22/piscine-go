package piscine

func ConcatParams(args []string) string {
	var str []string
	str = make([]string, len(args))
	for i, ch := range args {
		str[i] = ch
	}
	result := ""
	for i, r := range str {
		result += r
		if i < len(str)-1 {
			result += "\n"
		}
	}
	return result
}
