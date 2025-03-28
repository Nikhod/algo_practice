package main

import (
	"fmt"
)

func main() {
	fmt.Println("go algo")
	//nums2 := []int{2, 3, 4, 7, 11}
	//nums2 := []int{1, 2, 3, 4}
	nums2 := []int{8, 9, 12, 14, 16, 18, 25, 28, 30, 34, 37, 44, 50, 55, 58, 59, 73, 74, 81, 82, 83, 84, 87, 94, 96, 97,
		99, 103, 107, 115, 117, 119, 122, 123, 129, 130, 134, 137, 140, 142, 146, 149, 168, 170, 172, 174, 176, 178, 187,
		191, 197, 202, 203, 205, 207, 208, 211, 216, 218, 223, 228, 243, 244, 245, 257, 263, 267, 268, 271, 272, 274, 277,
		281, 285, 299, 303, 305, 314, 316, 321, 324, 327, 328, 329, 332, 333, 334, 363, 365, 367, 371, 373, 378, 380, 381,
		386, 389, 395, 400, 401, 414, 417, 425, 426, 427, 430, 431, 433, 434, 440, 441, 447, 449, 458, 459, 467, 468, 470,
		471, 472, 473, 474, 480, 482, 486, 487, 492, 502, 510, 515, 521, 525}
	//nums1 := []int{6, 13, 18, 18, 28, 34, 37, 39, 46, 50, 52, 54, 62, 63, 65, 66, 75, 80, 97, 98}
	//nums2 := []int{10, 13, 13, 19, 27, 33, 40, 41, 43, 46, 56, 61, 69, 72, 78, 79, 82, 88, 91, 94}
	fmt.Println(findKthPositive(nums2, 199))
}

func findKthPositive(arr []int, k int) int {
	ggg := 0
	// todo pivot даркор нест, достаточно обозначить верхнуюю границу и начиная с 1, проходиться в цикле
	upperBound := len()
	var result []int
	for i := 0; i < len(pivot) && len(result) != k+3; i++ {
		index := findTarget(arr, pivot[i])
		if index == -1 {
			result = append(result, pivot[i])
		}
		if i == 257 {
			ggg = 257
		}

	}

	//if len(result) > 0 {
	//	return result[k]
	//}
	fmt.Println(ggg)
	fmt.Println(result[199])
	fmt.Println(result)
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

func developArr() {
	var result []int
	for i := 0; i < 1001; i++ {
		result = append(result, i)
	}

	for _, v := range result {
		fmt.Print(v, ", ")
	}
	//fmt.Println(result)
}
