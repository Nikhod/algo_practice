package other_soulutions

func CountSort(list []int, maxValue int) {
	//maxValue = 4
	//list = []int{0, 2, 3, 0, 2, 0, 1, 0, 2, 3, 0, 2, 1,
	//	0, 1, 2, 0, 3, 0, 3, 2, 3, 2, 3, 1, 2, 1, 2, 1, 2}

	// value of 'list' == index of 'count'
	//sortedArray := make([]int, len(list), len(list))
	//count := make([]int, maxValue, maxValue)

	//min := list[0]
	//max := list[0]
	//for i := 0; i < len(list); i++ {
	//	if min < list[i] {
	//		min = list[i]
	//	}
	//	if list[i] < max {
	//		max = list[i]
	//	}
	//}

}

func CountingSort(list []int, maxValue int) []int {
	//maxValue = 4
	//list = []int{0, 2, 3, 0, 2, 0, 1, 0, 2, 3, 0, 2, 1,
	//	0, 1, 2, 0, 3, 0, 3, 2, 3, 2, 3, 1, 2, 1, 2, 1, 2}

	// index of 'count' == 'value' of list 		(list[i] = 'i' in loop)
	//sortedArray = [0,0,0,0,0,0,0,0,1,1,1,1,1,1,,2,2,2,2,2,2,2,3,3,3,3,3,]

	//sortedArray := make([]int, len(list), len(list))
	var sortedArray []int
	count := make([]int, maxValue, maxValue)

	for i := 0; i < len(list); i++ {
		count[list[i]]++
	}

	// i (index of count array) == value, which has to be in sorted array
	// в срезе 'count' хранятся количество каждого числа (его же индекс), сколько каждого элемента
	// value - сколько раз нужно добавить i (само число - значение) в
	for i, value := range count {
		for c := 0; c < value; c++ {
			sortedArray = append(sortedArray, i)
		}
	}

	return sortedArray
}
