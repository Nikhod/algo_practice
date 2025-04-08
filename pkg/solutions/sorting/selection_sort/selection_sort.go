package selection_sort

func Sort(arr []int) []int {
	var resultArr []int

	counter := len(arr)
	for i := 0; i < counter; i++ {
		smallestIndex := findSmallestValueIndex(arr)
		smallestValue := arr[smallestIndex]
		resultArr = append(resultArr, smallestValue)

		arr = removeValueFromArr(arr, smallestIndex)
	}

	return resultArr
}

func removeValueFromArr(arr []int, index int) []int {
	var result []int
	result = append(result, arr[:index]...)

	result = append(result, arr[index+1:]...)

	return result
}

func findSmallestValueIndex(arr []int) int {
	smallestNum := arr[0]
	smallestIndex := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < smallestNum {
			smallestNum = arr[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}
