package recursion

import "fmt"

func FindAmountOfNumberEntry() {
	arr := []int{23, 77, 44, 77, 56, 666, 78, 77, 88, 9999, 777, 77}
	usage := FindUsage(arr, 77, 0, 0)
	fmt.Println(usage)
}

func FindUsage(arr []int, number, index, counter int) int {

	if index >= len(arr) {
		return counter
	}

	if number == arr[index] {
		counter++
		return FindUsage(arr, number, index+1, counter)
	}

	return FindUsage(arr, number, index+1, counter)
}
