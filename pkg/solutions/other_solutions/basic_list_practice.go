package other_solutions

import (
	"container/list"
	"fmt"
)

func ListPractice() {
	l := list.New()
	//insert to front
	l.PushFront(1111)
	l.PushFront(2222)

	// insert to back
	l.PushBack(3333)
	l.PushBack(4444)

	// take an ELEMENT: back and front
	backElement := l.Back()
	frontElement := l.Front()

	//	l.Front() – возвращает первый элемент списка (начало).
	//	l.Back() – возвращает последний элемент списка (конец).

	// take a value from element
	backValue := backElement.Value
	frontValue := frontElement.Value

	fmt.Printf("back: %v\nfront: %v\n", backValue, frontValue)

	// iteration beginning from start
	for e := l.Front(); e != nil; e.Next() {
		value := e.Value
		fmt.Println(value)
	}

	// iteration beginning from end
	for e := l.Back(); e != nil; e.Prev() {
		value := e.Value
		fmt.Println(value)
	}
}
