package piscine

func SplitWhiteSpaces(s string) []string {
	result := make([]string, 0, len(s))
	word := ""
	for _, ch := range s {
		if ch == ' ' || ch == '\n' || ch == '\t' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
			continue
		}
		word += string(ch)
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
