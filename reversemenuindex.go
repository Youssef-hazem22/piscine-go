package piscine

func ReverseMenuIndex(menu []string) []string {
	youssef := make([]string, len(menu))
	for i := range menu {
		youssef[len(menu)-1-i] = menu[i]
	}
	return youssef
}
