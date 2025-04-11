package binary_search

// leetcode: 744
// You are given an array of characters letters that is sorted in non-decreasing order, and a character target.
//There are at least two different characters in letters.
//
//Return the smallest character in letters that is lexicographically greater than target. If such a character does
//not exist, return the first character in letters.
//
//Example 1:
//
//Input: letters = ["c","f","j"], target = "a"
//Output: "c"
//Explanation: The smallest character that is lexicographically greater than 'a' in letters is 'c'.
//Example 2:
//
//Input: letters = ["c","f","j"], target = "c"
//Output: "f"
//Explanation: The smallest character that is lexicographically greater than 'c' in letters is 'f'.
//Example 3:
//
//Input: letters = ["x","x","y","y"], target = "z"
//Output: "x"
//Explanation: There are no characters in letters that is lexicographically greater than 'z' so we return letters[0]

func NextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1

	for left <= right {
		mid := (right + left) / 2
		if letters[mid] == target {
			lastEntryIndex := findLastEntry(letters, letters[mid])

			if lastEntryIndex+1 < len(letters) {
				return letters[lastEntryIndex+1]
			} else {
				return letters[0]
			}

		}

		if letters[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= len(letters) {
		return letters[0]
	} else {
		return letters[left]
	}
}

func findLastEntry(arr []byte, target byte) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (right + left) / 2
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
