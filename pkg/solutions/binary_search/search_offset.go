package binary_search

// return index of target in list
func SearchOffset(list []int, target int) int {
	// data for testing
	//list := []int{4, 5, 6, 7, 0, 1, 2, 3}
	//target = 0
	// output := 4

	left := 0
	right := len(list) - 1

	for left <= right {
		mid := (left + right) / 2

		if list[mid] == target {
			return mid
		}

		// проверяем в какой части находится TARGET
		if list[left] <= list[mid] {
			//если он находится в левой части, то мы смещаем правую часть влево и присваиваем ее к mid
			if list[left] <= target && target < list[mid] {
				right = mid - 1
			} else { // если же target не находится в левой части, то мы смещаем левую границу вправо
				left = mid + 1
			}

		} else { // проверяем правую часть
			//если target находится в правой части (относительно mid)
			// то мы смещаем лувую границу вправо к mid
			if list[mid] < target && target <= list[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
