package recursion

//Вам даны два целочисленных массива nums1и nums2, отсортированных в неубывающем порядке , и два целых числа mи n,
//представляющие количество элементов в nums1и nums2соответственно.
//
//Объединить в один массив, отсортированный в nums1неубывающем порядке .nums2
//
//Окончательный отсортированный массив не должен возвращаться функцией, а должен храниться внутри массива nums1 .
//Чтобы учесть это, nums1имеет длину m + n, где первые m элементы обозначают элементы, которые должны быть объединены,
//а последние n элементы установлены в 0и должны игнорироваться. nums2имеет длину n.
//

// Пример 1:
//
// Вход: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Выход: [1,2,2,3,5,6]
// Пояснение: Массивы, которые мы объединяем, — это [1,2,3] и [2,5,6].
// Результатом слияния является [ 1 , 2,2 , 3,5,6 ] с подчеркнутыми элементами, полученными из nums1.
// Пример 2:
//
// Ввод: nums1 = [1], m = 1, nums2 = [], n = 0
// Вывод: [1]
// Пояснение: Массивы, которые мы объединяем, — это [1] и [].
// Результатом слияния является [1].
func merge(left []int, m int, right []int, n int) []int {
	// это для 2 примера
	if m == 0 && n > 0 {
		left[n-1] = right[n-1]
		n--
		return merge(left, m, right, n)
	}

	//basic recursion case
	if m-1 < 0 || n-1 < 0 {
		return left
	}

	k := m + n
	// будем сортировать с конца, поскольку нам нужно вернуть не новый срез, а старй left
	if left[m-1] > right[n-1] {
		left[k-1] = left[m-1]
		m--
		return merge(left, m, right, n)
	} else {
		left[k-1] = right[n-1]
		n--
		return merge(left, m, right, n)
	}
}
