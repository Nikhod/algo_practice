package quick_sort

func Sort(inputArr []int) []int {
	if len(inputArr) <= 1 {
		return inputArr
	} else {
		//	sequence:
		//	less -> pivot -> greater
		pivot := inputArr[0]
		less, greater := divideSlice(inputArr, pivot)

		less = Sort(less)
		greater = Sort(greater)

		less = append(less, pivot)
		less = append(less, greater...)
		return less
	}
}

// divideSlice: divide the slice to less than pivot and greater than pivot
func divideSlice(inputArr []int, pivot int) (less, greater []int) {
	for i := 1; i < len(inputArr); i++ {
		iterable := inputArr[i]
		if iterable > pivot {
			greater = append(greater, iterable)
		} else {
			less = append(less, iterable)
		}
	}
	return
}
