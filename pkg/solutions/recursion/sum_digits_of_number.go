package recursion

func FindSumOfDigits(num int) int {
	if num/10 == 0 {
		return num % 10
	}

	lastDigit := num % 10

	return lastDigit + FindSumOfDigits(num/10)
}
