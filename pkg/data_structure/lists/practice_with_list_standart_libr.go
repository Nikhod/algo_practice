package lists

import (
	"container/list"
	"fmt"
)

func printList(lst *list.List, msg string) {
	//lst.Front() - getting the first element in list
	//lst.Back()  - getting the last element in list
	fmt.Println(msg)
	for elm := lst.Front(); elm != nil; elm = elm.Next() {
		fmt.Printf("value: %v\t", elm.Value)
	}
	fmt.Println()
	fmt.Println()
	// ________________________________________________________________________________
	// another way of printing using the length of list
	//elm := lst.Front()
	//for i := 0; i < lst.Len(); i++ {
	//	fmt.Printf("value: %v\t", elm.Value)
	//	elm = elm.Next()
	//}
}

// push - adding to beginning of list
func PushingData() {
	firstList := list.New()
	secondList := list.New()

	secondList.PushFront(99999)
	secondList.PushFront(88888)

	//adding elements to list
	// every list element has index, but those are can be changed by moving/inserting/pushing new data
	firstList.PushBack(777)
	firstList.PushFront(111)
	firstList.PushFront(222)
	firstList.PushFront(333)
	printList(firstList, "BEFORE pushing second list")

	//adding the second list to the beginning of first list. After moving elements,
	//second list becomes empty, and first list length increase on len(secondList)
	firstList.PushFrontList(secondList)

	printList(firstList, "AFTER pushing second list")
}

func InsertingData() {
	firstList := list.New()

	frsElm := firstList.PushFront(111)
	firstList.PushFront(222)
	thirdElm := firstList.PushFront(333)

	// inserting data to list after third element
	firstList.InsertAfter(444444, thirdElm)
	firstList.InsertBefore(0000, frsElm)
	printList(firstList, "inserting after third, and before the first element")
}

func MovingData() {
	firstList := list.New()

	firstList.PushFront(111)
	secondElement := firstList.PushFront(222)
	thirdElm := firstList.PushFront(333)
	printList(firstList, "after init the list: ")

	//moving element to the very start
	firstList.MoveToFront(secondElement)
	//moving element to the very end
	firstList.MoveToBack(thirdElm)

	printList(firstList, "after moving the start/end the second and third element respectively")
}

func RemovingData() {
	firstList := list.New()

	//adding elements to list
	last := firstList.PushBack(777)
	firstList.PushFront(111)
	firstList.PushFront(222)
	firstList.PushFront(333)
	printList(firstList, "after initialization")

	firstList.Remove(last)
	printList(firstList, "removing the last element - (777)")
}
