package binary_search

// leetcode: 1539
// Дан массив arr положительных целых чисел, отсортированный в строго возрастающем порядке , и целое число k.
// Верните положительное целое число, отсутствующее в этом массиве.kth
// Пример 1:
// Вход: arr = [2,3,4,7,11], k = 5
// Выход: 9
// Пояснение: Недостающие положительные целые числа: [1,5,6,8,9,10,12,13,...].
// Пятое недостающее  положительное целое число: 9.
func findKthPositive(arr []int, k int) int {
	var result []int
	var i = 1
	for len(result) != k+1 {
		index := findTarget(arr, i)
		if index == -1 {
			result = append(result, i)
			i++
			continue
		}
		if len(result) == k+2 {
			break
		}
		i++
	}

	//if len(result) > 0 {
	//	return result[k]
	//}

	return result[k-1]
}

func findTarget(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == target {
			return mid
		}
		if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
