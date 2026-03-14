package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var f food

	if order == "burger" {
		f.preptime = 15
	} else if order == "chips" {
		f.preptime = 10
	} else if order == "nuggets" {
		f.preptime = 12
	} else {
		return 404
	}

	return f.preptime
}
