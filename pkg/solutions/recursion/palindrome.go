package recursion

// one can optimize this func by using the  array of rune or byte (uint8)
func IsPalindrome(text []byte) bool {
	if len(text) == 1 || len(text) == 0 {
		return true
	}
	if text[0] == text[len(text)-1] {
		text = text[1 : len(text)-1]
		return IsPalindrome(text)
	}

	return false
}
