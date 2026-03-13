package piscine

func ToLower(s string) string {
	S := ""
	for _, ch := range s {
		if !(ch >= 65 && ch <= 90) {
			S += string(ch)
		} else {
			ch += 32
			S += string(ch)
		}
	}
	return S
}
