package quad

func QuadA(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res string
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if (row == 1 || row == y) && (col == 1 || col == x) {
				res += "o"
			} else if row == 1 || row == y {
				res += "-"
			} else if col == 1 || col == x {
				res += "|"
			} else {
				res += " "
			}
		}
		res += "\n"
	}
	return res
}

func QuadB(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res string
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				res += "/"
			} else if row == 0 && col == x-1 {
				res += "\\"
			} else if row == y-1 && col == 0 {
				res += "\\"
			} else if row == y-1 && col == x-1 {
				res += "/"
			} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
				res += "*"
			} else {
				res += " "
			}
		}
		res += "\n"
	}
	return res
}

func QuadC(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res string
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 {
				if col == 1 || col == x {
					res += "A"
				} else {
					res += "B"
				}
			} else if row == y {
				if col == 1 || col == x {
					res += "C"
				} else {
					res += "B"
				}
			} else {
				if col == 1 || col == x {
					res += "B"
				} else {
					res += " "
				}
			}
		}
		res += "\n"
	}
	return res
}

func QuadD(x, y int) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	var res string
	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if (row == 0 || row == y-1) && col == 0 {
				res += "A"
			} else if (row == 0 || row == y-1) && col == x-1 {
				res += "C"
			} else if row == 0 || row == y-1 {
				res += "B"
			} else if col == 0 || col == x-1 {
				res += "B"
			} else {
				res += " "
			}
		}
		res += "\n"
	}
	return res
}

func QuadE(c, r int) string {
	if c <= 0 || r <= 0 {
		return ""
	}
	var res string
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if i == 0 && j == 0 || (i == r-1 && j == c-1 && !(r == 1 || c == 1)) {
				res += "A"
			} else if i == r-1 && j == 0 || (i == 0 && j == c-1) {
				res += "C"
			} else if (i == 0 && j > 0 && j < c-1) || (i == r-1 && j > 0 && j < c-1) {
				res += "B"
			} else if (j == 0 && i > 0 && i < r-1) || (j == c-1 && i > 0 && i < r-1) {
				res += "B"
			} else {
				res += " "
			}
		}
		res += "\n"
	}
	return res
}
