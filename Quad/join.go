package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i := 0; i < len(strs); i++ {
		if i > 0 {
			result += sep
		}
		result += strs[i]
	}
	return result
}
