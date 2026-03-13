package piscine

func Capitalize(s string) string {
	result := []rune(s)
	con := true
	for i, ch := range result {
		if !((ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')) {
			con = true
			continue
		}
		if con {
			if ch >= 'a' && ch <= 'z' {
				result[i] -= 32
			}
			con = false
		} else {
			if ch >= 'A' && ch <= 'Z' {
				result[i] += 32
			}
		}
	}
	return string(result)
}
