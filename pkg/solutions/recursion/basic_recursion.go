package recursion

import "fmt"

// BasicRecursion : countdown beginning from 'number'0
func BasicRecursion(number int) {
	fmt.Println(number)
	if number == 0 {
		return
	}
	BasicRecursion(number - 1)
}
