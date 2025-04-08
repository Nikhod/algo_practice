package merge_sort

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	sorted := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] { // изменив знак на обратный, можно отсортировать в убывание
			sorted = append(sorted, left[l])
			l++
		} else {
			sorted = append(sorted, right[r])
			r++
		}
	}
	// переменные l & r после цикла могут быть не равны из за разной длины срезов, поэтому один из них будет
	// обозначать еще индекс элемента который не был учтен в обработке и сравнении (это будет 1 элемент),
	// и это будет самое большое число, он пойдет в конец отсортированного среза.
	sorted = append(sorted, left[l:]...)
	sorted = append(sorted, right[r:]...)

	return sorted
}
