package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""
	for _, ch := range str {
		if ch == ' ' || ch == '\n' || ch == '\t' {
			summary[word]++
			word = ""
		} else {
			word += string(ch)
		}
	}
	summary[word]++
	return summary
}
