package recursion

// leetcode: 3304
// Алиса и Боб играют в игру. Изначально у Алисы есть строка word = "a".
// Вам дано положительное целое число k.
// Теперь Боб попросит Алису выполнить следующую операцию навсегда :
// Создайте новую строку, заменив каждый символ wordна следующий символ английского алфавита,
// и добавьте ее к исходной word .
// Например, выполнение операции над "c"generates "cd"и выполнение операции над "zb"generates "zbac".
// Возвращает значение символа в , после того как будет выполнено достаточное количество
// операций для получения по крайней мере символов.kthwordword k
// Обратите внимание , что в ходе операции 'z'можно изменить символ .'a'
func KthCharacter(k int) byte {
	word := []byte{97}
	ans := delay(k, word)
	return ans
}

func delay(k int, word []byte) byte {
	if k < len(word) {
		return word[k-1]
	}

	sufWord := findSuf(word)
	word = append(word, sufWord...)

	return delay(k, word)
}

func findSuf(word []byte) []byte {
	result := []byte{}
	for i := 0; i < len(word); i++ {
		if word[i] == 122 {
			result = append(result, 97)
		}

		result = append(result, word[i]+1)
	}

	return result
}
