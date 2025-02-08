package other_solutions

import "strings"

func IsPalindrome(text string) bool {
	text = strings.ToLower(text)

	text = RemoveSpace(text)
	reverseCounter := len(text) - 1

	for i := 0; i < len(text); i++ {
		if text[i] != text[reverseCounter] {
			return false
		}
		reverseCounter--
	}

	return true
}

func RemoveSpace(text string) string {
	var clearSlice []uint8
	for i := 0; i < len(text); i++ {
		if char := text[i]; rune(char) != ' ' {
			clearSlice = append(clearSlice, char)
		}
	}

	return string(clearSlice)
}
