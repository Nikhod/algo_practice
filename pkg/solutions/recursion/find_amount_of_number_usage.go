package recursion

import "fmt"

func FindAmountOfNumberEntry() {
	arr := []int{23, 77, 44, 77, 56, 666, 78, 77, 88, 9999, 777, 77}
	usage := FindUsage(arr, 77, 0)
	fmt.Println(usage)
}

func FindUsage(arr []int, number, index int) int {
	var counter int

	if index >= len(arr) {
		return -1
	}

	if number == arr[index] {
		return counter
	}

	return FindUsage(arr, number, index+1)
}
