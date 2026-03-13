package piscine

func Rot14(s string) string {
	youssef := ""
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			youssef += string('a' + (ch-'a'+14)%26)
		} else if ch >= 'A' && ch <= 'Z' {
			youssef += string('A' + (ch-'A'+14)%26)
		} else {
			youssef += string(ch)
		}
	}
	return youssef
}
