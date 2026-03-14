package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1

	for _, ch := range s {
		if ch == '-' && num == 0 {
			sign = -1
		} else if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
		}
	}
	return num * sign
}
