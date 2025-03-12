package binary_search

// Дан массив nums, который почти отсортирован:
//
//	Каждый элемент находится на своём месте или сдвинут максимум на одну позицию влево или вправо.
//	Тебе нужно найти индекс элемента target.
//
// nums := []int{10, 3, 40, 20, 50, 80, 70}
// target := 40
func FindTargetWithOffsetByOneElement(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		// сдвиг влево
		if left < mid && nums[mid-1] == target {
			return mid - 1
		}
		// сдвиг вправо
		if mid < right && nums[mid+1] == target {
			return mid + 1
		}

		// изменение границ
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
