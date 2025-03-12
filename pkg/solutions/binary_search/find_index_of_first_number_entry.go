package binary_search

// Условие:
// Дан отсортированный массив целых чисел nums, который может содержать повторяющиеся элементы. Нужно найти индекс первого вхождения числа target. Если target отсутствует, вернуть -1.
//
// Пример 1:
//
// go
// Копировать
// Редактировать
// nums := []int{1, 2, 2, 2, 3, 4, 5}
// target := 2
func FindFirstEntry(list []int, target int) int {
	left := 0
	right := len(list) - 1
	resultIndex := -1
	for left < right {
		mid := (right + left) / 2
		if list[mid] == target {
			resultIndex = mid
			right = mid - 1
			continue
		}

		if list[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return resultIndex
}
