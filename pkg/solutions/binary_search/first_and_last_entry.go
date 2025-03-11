package binary_search

import (
	"fmt"
)

// Поиск первой и последней позиции элемента в отсортированном массиве
// Дан отсортированный массив целых чисел nums и целое число target. Найдите первую и последнюю позицию target в
// массиве. Если target не присутствует в nums, верните [-1, -1].
// Твое решение должно работать за O(log n).
// the only case, where the solution work incorrectly
// nums := []int{5, 7, 7, 8, 8, 10}
// target := 8
func FindFirstAndLastEntry(nums []int, target int) (int, int) {
	first, last := -1, -1

	if len(nums) == 0 {
		return -1, -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0, 0
		}
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			if first == -1 {
				first = mid
			} else {
				last = mid
			}
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if first != -1 && last == -1 {
		last = first
	} else {
		first = last
	}
	return first, last
}

// the only case, where the solution work incorrectly
// nums := []int{5, 7, 7, 8, 8, 10}
// target := 8
func SecondFindFirstAndLastEntry(nums []int, target int) (int, int) {
	var entries []int
	if len(nums) == 0 {
		return -1, -1
	}

	if len(nums) == 1 {
		if nums[0] == target {
			return 0, 0
		}
	}

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			entries = append(entries, mid)
			left = mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println(entries)
	if len(entries) == 0 {
		return -1, -1
	}
	return entries[0], entries[len(entries)-1]
}

// nums := []int{5, 7, 7, 8, 8, 10}
// target := 8
func ThirdFindFirstAndLastEntry(nums []int, target int) (int, int) {
	// Найти первое вхождение
	first, last := -1, -1
	left, right := 0, len(nums)-1
	for left <= right {
		// ищем первое вхождение: перемещаем правую границу влево, если мы нашли target, в противном случае
		//перемещаем левую границу вправо
		mid := (right + left) / 2
		if target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// есть ли target в массиве
	if left < len(nums) {
		if target == nums[left] {
			first = left
		} else {
			return first, last
		}
	}

	left, right = 0, len(nums)-1
	// ищем последнее вхождение: перемещаем левую границу вправо если нашли target
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	last = right

	return first, last
}
