package counting_sort

import "fmt"

// в этом подходе мы создаем два массива и потом заполняем их числами сколько раз каждый символ
// встречался в каждой строке, соответственно, получается, что один массив записывает количество букв, которые были
// встречены в target, а другой в checkString
// После того, как два массива были заполнены, мы сравниваем в конце этих два массива. И смотрим, были же те же
// самые буквы в Target и в checkString
func IsAnagram(checkString, target string) bool {
	//data for testing
	//checkString := "kufdfadfkyuuukk"
	//target := "kuku"
	//output: true

	//В задачах с отступом нужно учитывать еще и минимальное (97) и максимальное значение (122) какого-то массива
	//чтобы найти их разницу. Разница = 26 => создаем массив\срез с объемом 26
	maxValue := 26 // there are 26 char in latin alphabet
	countCheckString := make([]int, maxValue, maxValue)
	countTarget := make([]int, maxValue, maxValue)

	// the alphabet starts by char 'a' and ends by 'z', whose index number is 97 and 122 respectively
	// therefore we develop the int slice, whose indexes will be meant index number of char, but index
	// 0 means char 'a', not 97 as usual
	// It will optimize memory

	//countCheckString = CountString(checkString, countCheckString)
	countTarget = CountString(target, countTarget)

	// optimization: можно i вынести как counter, а условие поставить так, чтобы итерации были до тех пор пока кол-во
	// вхождений буквы в countTarget и countCheckString не будут равны, и сделать выход из цикла через counter,
	// если он равен len(checkString). После цикла return false, это будет означать что букв в checkString не хватает
	// чтобы удовлетворить target
	for i := 0; i < len(checkString); i++ {
		byteFromTarget := checkString[i]
		indexInNewSlice := byteFromTarget - 97
		if byteFromTarget == 32 {
			continue
		}

		//добавляем в срез только те символы, которые есть в target (countTarget)
		if countTarget[indexInNewSlice] > 0 {
			countCheckString[indexInNewSlice]++
		}
	}

	var result bool = true
	for i := 0; i < len(countTarget); i++ {
		// если количество буквы вошедшей в checkString >= в target (необходимый минимум) -> true
		// если этот минимум не соблюдается то -> false
		if value := countCheckString[i]; value < countTarget[i] {
			result = false
		}
	}

	fmt.Println(countCheckString)
	fmt.Println(countTarget)

	return result
}

// сколько каких букв в строке
func CountString(target string, countTarget []int) []int {
	for i := 0; i < len(target); i++ {
		byteFromTarget := target[i]
		if byteFromTarget == 32 {
			continue
		}
		//Так как в таблице символов ‘a’ начинается с 97 мы должны
		//получить порядковый номер символа в строке и отнять 97
		countTarget[byteFromTarget-97]++
	}

	return countTarget
}
