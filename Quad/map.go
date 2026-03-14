package piscine

func isprime(b int) bool {
	if b <= 1 {
		return false
	}
	if b == 2 {
		return true
	}
	if b%2 == 0 {
		return false
	}
	for i := 2; i <= b-1; i++ {
		if b%i == 0 {
			return false
		}
	}
	return true
}

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, ch := range a {
		result[i] = f(ch)
	}
	return result
}
