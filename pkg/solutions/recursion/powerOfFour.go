package recursion

// leetcode task
// является ли число степенью 4
func isPowerOfFour(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 {
		return false
	}

	if n%4 == 0 {
		return isPowerOfFour(n / 4)
	}
	return false
}
