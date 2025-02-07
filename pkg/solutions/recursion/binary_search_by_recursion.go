package recursion

func FindNumber(arr []int, item int) int {
	if len(arr) == 1 {
		return arr[1]
	}

	left := 0
	right := len(arr) - 1
	mid := (left + right) / 2

	if arr[mid] == item {
		return arr[mid]
	}

	if arr[mid] < item {
		return FindNumber(arr[mid+1:], item)
	} else {
		return FindNumber(arr[:mid], item)
	}
}
