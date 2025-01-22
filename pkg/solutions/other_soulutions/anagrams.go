package other_soulutions

import (
	"unicode"
)

// строка A нужно переставить буквы так, чтобы получилась строка Б
func IsAnagram(anagram, target string) bool {
	//data for testing
	//anaram := "kufdfadfkyuuukk"
	//target := "kuku"

	// Создаем карту для символов второй строки
	mapTarget := make(map[rune]int)
	makingTargetMapWithRuneChar(target, mapTarget)

	// Уменьшаем счетчики символов для каждого символа target если нашли такую же в anagram
	for _, runeChar := range anagram {

		//ингорируем пробелы и лишние символы
		if _, ok := mapTarget[runeChar]; ok == false || unicode.IsSpace(runeChar) {
			continue
		}
		// декрименция счетчика
		if _, ok := mapTarget[runeChar]; ok {
			mapTarget[runeChar]--
		}
	}

	// Проверяем, что все значения в карте равны 0, если нет, то значит какого то символа нет в anagram
	// чтобы состваить из нее target. Счетчик может быть отрицательным, что значит избыток некторых символов
	// в anagram
	for _, value := range mapTarget {
		if value > 0 {
			return false
		}
	}

	return true
}

// превращает строку в hash-map чтобы посчитать сколько раз встречается каждый символ
func makingTargetMapWithRuneChar(target string, mapTarget map[rune]int) {
	for i := 0; i < len(target); i++ {
		runeCharTarget := rune(target[i])

		if unicode.IsSpace(runeCharTarget) {
			continue
		}

		mapTarget[runeCharTarget]++
	}
}
