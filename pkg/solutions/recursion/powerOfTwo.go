package recursion

func isPowerOfTwo(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 {
		return false
	}

	// если число делится на 2, то продолжаем его делить и дойдем до true
	if n%2 == 0 {
		return isPowerOfTwo(n / 2)
	}
	return false
}
