package two_pointers

// leetCode: 2540
// Даны два целочисленных массива nums1и nums2, отсортированные в неубывающем порядке, вернуть минимальное целое число,
// общее для обоих массивов nums1 . Если общего целого числа среди и нет nums2, вернуть -1.
// Обратите внимание, что целое число считается общим для nums1, nums2если в обоих массивах есть хотя бы одно вхождение
// этого целого числа.
func GetCommon(nums1 []int, nums2 []int) int {
	minimum := -1

	if len(nums1) == 0 || len(nums2) == 0 {
		return -1
	}

	var firstPointer, secondPointer int = 0, 0
	for firstPointer < len(nums1) && secondPointer < len(nums2) {
		if nums1[firstPointer] == nums2[secondPointer] {
			minimum = nums1[firstPointer]
			return minimum
		} else if nums1[firstPointer] < nums2[secondPointer] {
			firstPointer++
		} else if nums1[firstPointer] > nums2[secondPointer] {
			secondPointer++
		}
	}
	return minimum
}
