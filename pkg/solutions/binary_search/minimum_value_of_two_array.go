package binary_search

// leetCode: 2540
// Даны два целочисленных массива nums1и nums2, отсортированные в неубывающем порядке, вернуть минимальное целое число,
// общее для обоих массивов nums1 . Если общего целого числа среди и нет nums2, вернуть -1.
// Обратите внимание, что целое число считается общим для nums1, nums2если в обоих массивах есть хотя бы одно вхождение
// этого целого числа.
func GetCommon(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return -1
	}
	var minimum int = -1
	for _, numInFirstArr := range nums1 {
		indexOfNumInFirstArr := findTargetByBinarySearch(nums2, numInFirstArr)
		if indexOfNumInFirstArr == -1 {
			continue
		} else {
			return nums2[indexOfNumInFirstArr]
		}
	}

	return minimum
}
func findTargetByBinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (right + left) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func getCommon(nums1 []int, nums2 []int) int {
	minimum := -1

	if len(nums1) == 0 || len(nums2) == 0 {
		return -1
	}

	var firstPointer, secondPointer int = 0, 0
	// все 2 указателя двигаются справа налево, поскольку надо найти минимальное значение
	for firstPointer < len(nums1) && secondPointer < len(nums2) {
		if nums1[firstPointer] == nums2[secondPointer] {
			minimum = nums1[firstPointer]
			return minimum
			// дальше, мы двигаем тот указатель,который указывает на число меньшее по значению. Так как нам нуж
			// нужно найти общее минимальное число, то если число под индексом firstPointer 1ого массива меньше числа
			// под индексом secondPointer, его (firstPointer) надо двигать вправо.
		} else if nums1[firstPointer] < nums2[secondPointer] {
			firstPointer++
		} else if nums1[firstPointer] > nums2[secondPointer] {
			secondPointer++
		}
	}
	return minimum
}
