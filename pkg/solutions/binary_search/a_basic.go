package binary_search

import (
	"fmt"
	"time"
)

func BinarySearch(list []int, item int) *int {
	//data for testing
	//list := []int{1, 3, 5, 6, 7, 9, 11, 13, 17, 21}
	//item := 11

	left := 0
	right := len(list) - 1

	var counter int
	for left <= right {
		mid := (left + right) / 2
		guess := list[mid]

		if guess == item {
			return &mid
		}

		if guess < item {
			left = mid + 1
		} else {
			right = mid - 1
		}

		fmt.Println("guess: ", guess, "iteration # ", counter+1)
		counter++
		time.Sleep(time.Second * 2)
	}
	return nil
}
