package recursion

//	this solution has one issue - the indexes during the executing the program is being lost
func FindNumber(arr []int, item int) int {
	// if the array is empty, return -1
	if len(arr) == 0 {
		return -1
	}

	// basic case [is very important]
	if len(arr) == 1 {
		return arr[0]
	}

	//	putting the edges
	left := 0
	right := len(arr) - 1
	mid := (left + right) / 2

	// case where the item is found
	if arr[mid] == item {
		return arr[mid]
	}

	//	consider the cases when item is situated in left or right sides of the array
	if arr[mid] < item {
		return FindNumber(arr[mid+1:], item)
	} else {
		return FindNumber(arr[:mid], item)
	}
}
