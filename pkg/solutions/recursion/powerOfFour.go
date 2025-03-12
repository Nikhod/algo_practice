package recursion

// leetcode task
func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 {
		return false
	}

	if n%4 == 0 {
		isPowerOfFour(n / 4)
	}
	return false
}
