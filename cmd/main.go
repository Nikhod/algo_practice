package main

import (
	"fmt"
)

func main() {
	fmt.Println("go algo")
	//nums1 := []int{1, 2, 3, 6}
	//nums2 := []int{2, 3, 4, 5}
	nums1 := []int{6, 13, 18, 18, 28, 34, 37, 39, 46, 50, 52, 54, 62, 63, 65, 66, 75, 80, 97, 98}
	nums2 := []int{10, 13, 13, 19, 27, 33, 40, 41, 43, 46, 56, 61, 69, 72, 78, 79, 82, 88, 91, 94}

	fmt.Println(getCommon(nums1, nums2))
}
