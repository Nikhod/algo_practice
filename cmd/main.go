package main

import (
	"algo/pkg/solutions/binary_search"
	"fmt"
)

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8

	first, last := binary_search.ThirdFindFirstAndLastEntry(nums, target)
	//first, last := FindFirstAndLastEntry(nums, target)
	fmt.Println(first, last)
}
