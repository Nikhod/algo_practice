package recursion

import "fmt"

func PracticeRecursion() {
	arr := []int{23, 44, 56, 666, 78, 88, 9999, 777}
	index := FindNumb(arr, 777, 0)
	fmt.Println("the value in the main func: ", index)
	fmt.Println("value of arr[index]: ", arr[index])
}

func FindNumb(arr []int, number, index int) int {
	if index >= len(arr) {
		return -1
	}

	if arr[index] == number {
		return index
	}

	return FindNumb(arr, number, index+1)
}
