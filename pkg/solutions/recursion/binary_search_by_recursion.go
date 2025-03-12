package recursion

// this solution has one issue - the indexes during the executing the program is being lost
func BinarySearchByRecursion(arr []int, target int, left, right int) int {
	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		if arr[0] == target {
			return 0
		} else {
			return -1
		}
	}

	mid := (right + left) / 2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] < target {
		return BinarySearchByRecursion(arr, target, mid+1, right)
	} else {
		return BinarySearchByRecursion(arr, target, left, mid)
	}
}
