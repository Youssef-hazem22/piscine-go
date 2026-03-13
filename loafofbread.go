package piscine

func LoafOfBread(str string) string {
	result := ""
	count := 0
	skipping := false

	for _, r := range str {
		if r == ' ' {
			continue
		}
		if skipping {
			skipping = false
			continue
		}
		result += string(r)
		count++
		if count == 5 {
			result += " "
			count = 0
			skipping = true
		}
	}

	if count == 0 && len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	if len(result) == 0 {
		return "Invalid Output\n"
	}

	return result + "\n"
}
