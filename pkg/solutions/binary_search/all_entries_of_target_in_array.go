package binary_search

import "fmt"

// leetcode: 2089
func TargetIndices(nums []int, target int) []int {
	nums = qSort(nums)
	var resultIndexes []int
	var first, last int
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right + left) / 2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	first = left
	fmt.Println("len:", len(nums))
	fmt.Println("first", first)

	left, right = 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//last = len(nums) - left
	last = left - 1
	fmt.Println("Last", last)

	for i := first; i <= last; i++ {
		resultIndexes = append(resultIndexes, i)
	}

	return qSort(resultIndexes)
}

func qSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	} else {

		var pivot int
		if len(nums) != 0 {
			pivot = nums[0]
		}
		less, greater := greaterAndLess(nums[1:], pivot)

		return append(append(qSort(less), pivot), qSort(greater)...)
	}
}

func greaterAndLess(nums []int, pivot int) ([]int, []int) {
	var less []int
	var greater []int
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			less = append(less, nums[i])
		} else {
			greater = append(greater, nums[i])
		}
	}
	return less, greater
}
