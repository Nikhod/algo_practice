package counting_sort

import (
	"fmt"
	"strings"
)

// FindAmountOfUsageOfSubstring :
//	обозначаем константу space, потому что 32 это байт, который указывает на порядковый номер в таблице ASCII - пробел
// Дальше мы приводим к общему знаменателю, т.е.
// Снижаем регистр до 1, потом мы инициализируем. массив или срез с размером количества букв в алфавите английском,
// то есть 26. И дальше находим порядковый номер каждого символа в алфабете.
// мы передаем функцию target строку
// и в этой таргет строке мы будем добавлять отдельный срез,
// И будем находить Порядковый номер в алфавите тех символов, которые встретились в Таргет.
// дальше идет у нас цикл который пропускает все пробелы и считает сколько каждый раз какой-то символ в большой строке
// встретился
// И нам нужно свериться и найти минимальное значение Сколько раз буква, которая удовлетворяет таргету, была вхождена в
// большую строку.
func FindAmountOfUsageOfSubstring(largeStr string, target string) {
	const SPACE uint8 = 32
	largeStr = strings.ToLower(largeStr)
	count := make([]int, 26, 26)
	positionOfEachTargetLetter := findCharPosition(target) // находим порядковый номер каждого символа в алфавите

	for i := 0; i < len(largeStr); i++ {
		if largeStr[i] == SPACE {
			continue
		}
		count[largeStr[i]-97]++
	}
	fmt.Println(count)
	fmt.Println(positionOfEachTargetLetter)

	min := count[0]
	for i := 0; i < len(positionOfEachTargetLetter); i++ {
		amountOfCountLettersUsages := count[positionOfEachTargetLetter[i]]
		if amountOfCountLettersUsages < min {
			min = amountOfCountLettersUsages
		}
	}

}

//	Число девяносто семь встречается здесь два раза, оба раза А неодинаковые случаи на самом деле. Почему 97?
//	Потому что в... таблицы Unicode и ASCII, буква A встречается на девяносто седьмом порядке. Буква Z, то есть
//	максимальное значение. встречается под порядковым номером 122.
//	Поэтому минимальное значение 97 из латинского алфавита. И все, что в промежутке, это все те символы, которые будут
//	встречаться. В том числе 97 и 122 Поэтому Каждый руну мы должны отнять 97 чтобы получить порядковый номер в срезе count
func findCharPosition(target string) []uint8 {
	var indexListOfEachLetter []uint8
	const FirstPositionOfFirstLetter uint8 = 97

	for i := 0; i < len(target); i++ {
		char := target[i]
		position := char - FirstPositionOfFirstLetter
		indexListOfEachLetter = append(indexListOfEachLetter, position)
	}

	return indexListOfEachLetter
}
