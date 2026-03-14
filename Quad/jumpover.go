package piscine

func JumpOver(str string) string {
	if len(str) < 3 {
		return "\n"
	}
	counter := 0
	word := ""
	for i := 0; i < len(str); i++ {
		counter++
		if counter%3 == 0 && i != 0 {
			word += string(str[i])
		}
	}
	return word + "\n"
}
