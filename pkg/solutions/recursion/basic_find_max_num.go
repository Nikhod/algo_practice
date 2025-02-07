package recursion

// algo - method "divide and conquer" is used
func FindMaxInList(arr []int) int {
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		} else {
			return arr[1]
		}
	}

	sub_max := FindMaxInList(arr[1:])
	if arr[0] > sub_max {
		return arr[0]
	} else {
		return sub_max
	}
}
