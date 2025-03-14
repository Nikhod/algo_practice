package binary_search

func MaximumCount(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2

		if nums[mid] < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	neg := left

	left, right = 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] <= 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	pos := len(nums) - left

	if neg > pos {
		return neg
	}
	return pos
}
