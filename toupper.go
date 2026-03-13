package piscine

func ToUpper(s string) string {
	S := ""
	for _, ch := range s {
		if !(ch >= 97 && ch <= 122) {
			S += string(ch)
		}
		if ch >= 97 && ch <= 122 {
			ch -= 32
			S += string(ch)
		}
	}
	return S
}
