package recursion

// find the sum of number in array by algo - method "divide and conquer"
func SumList(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + SumList(arr[1:])
}
