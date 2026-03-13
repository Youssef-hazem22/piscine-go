package piscine

func Split(s, sep string) []string {
	result := ""
	parts := make([]string, 0, len(s))

	i := 0
	for i < len(s) {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			if result != "" {
				parts = append(parts, result)
				result = ""
			}
			i += len(sep)
		} else {
			result += string(s[i])
			i++
		}
	}

	if result != "" {
		parts = append(parts, result)
	}

	return parts
}
