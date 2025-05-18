package quick_sort

import "math/rand"

func SortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	} else {
		//	sequence:
		//	less -> pivot -> greater
		pivotIndex := rand.Intn(len(nums) - 1)
		pivot := nums[pivotIndex]
		less, greater := divideSlice(nums, pivot, pivotIndex)

		less = SortArray(less)
		greater = SortArray(greater)

		less = append(less, pivot)
		less = append(less, greater...)
		return less
	}
}

// divideSlice: divide the slice to less than pivot and greater than pivot
func divideSlice(inputArr []int, pivot, pivotIndex int) (less, greater []int) {
	// pivotIndex нужен чтобы мы дважды не обработали сам опорный элемент, поскольку он обрабатывается в
	// другой функции
	for i := 0; i < len(inputArr); i++ {
		if i == pivotIndex {
			continue
		}
		iterable := inputArr[i]
		if iterable > pivot {
			greater = append(greater, iterable)
		} else {
			less = append(less, iterable)
		}
	}
	return
}
func Swap(nums []int) ([]int, int) {
	pivotIndex := (len(nums) - 1) / 2
	pivot := nums[pivotIndex]

	left := nums[:pivotIndex]
	right := nums[pivotIndex:]
	left = append(left, right...)
	nums = left

	var l = -1
	for g := 0; g < len(nums); g++ {
		if nums[g] <= pivot {
			l++
			nums[l], nums[g] = nums[g], nums[l]
		}
	}
	//l - индекс-граница между большими и меньшими чем опорный
	return nums, l
}
