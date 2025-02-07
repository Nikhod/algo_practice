package recursion

// find amount of element in array
func CountAmountOfElement(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return 1 + CountAmountOfElement(arr[1:])
}
