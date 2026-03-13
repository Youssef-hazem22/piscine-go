package piscine

func DivMod(a int, b int, div *int, mod *int) {
	x := a / b
	y := a % b
	*div = x
	*mod = y
}
