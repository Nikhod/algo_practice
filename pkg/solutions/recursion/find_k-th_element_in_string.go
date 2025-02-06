package recursion

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
