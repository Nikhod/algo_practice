package lists

// push - set value
// front - first\beginning of list
// back - last\ending of list

func New() *LinkedList {
	return &LinkedList{head: nil}
}

type Node struct {
	value any
	next  *Node
}

type LinkedList struct {
	head *Node
}

// todo check all the topic

func (l *LinkedList) RemoveBack() {
	if l.head == nil {
		return
	}
	// finding the last not empty node
	current := l.head
	for {
		// если в цикле мы увидим что следующий элемент последний в списке, мы выходим из списка
		// и поле next предпоследнего элемента становится nil, то есть становится последним
		if current.next == nil {
			break
		}
		current = current.next

		// "current = current.next" != current = l.head.next. Одна ячейка памяти но 2 разные переменные
		// каждый "current = l.head.next" - каждый раз будет обращаться ко второму элементу, поскольку
		// l.head - первый элемент в списке.
		//
		//
		//
		//current := l.head
		//Теперь:
		//
		//l.head → *Node (допустим, адрес 0x123)
		//
		//current → та же *Node (0x123) — копия указателя
		//current = current.next
		//Ты не изменяешь l.head — ты просто меняешь current, чтобы он указывал на другой узел (например, 0x456),
		//но l.head всё ещё указывает на 0x123.
		// ❗ Но если ты меняешь сам объект, на который указывает current, тогда head тоже "почувствует" это изменение:
		//current := head
		//current.value = 99
		//Теперь:
		//
		//head.value == 99 ✅ потому что current и head указывают на один и тот же объект
	}
	current.next = nil
}

func (l *LinkedList) RemoveFront() {
	if l.head == nil {
		return
	}
	l.head = l.head.next
}

func (l *LinkedList) PushFront(value any) {
	// if there is no element in list
	if l.head == nil {
		newNode := Node{
			value: value,
			next:  nil,
		}
		l.head = &newNode
		return
	}

	newNode := Node{
		value: value,
		next:  l.head,
	}
	// current is the valuable to be moved after newNode
	l.head = &newNode
}

func (l *LinkedList) PushBack(value any) {
	newNode := Node{
		value: value,
		next:  nil,
	}
	if l.head.next == nil {
		l.head = &newNode
		return
	}

	//find current, reach last element (with next = nil)
	current := l.head
	for current != nil {
		current = current.next
	}

	current = &newNode
	return
}

func (l *LinkedList) PushB(value any) {
	if l.head == nil {
		newNode := Node{
			value: value,
			next:  nil,
		}
		l.head = &newNode
		return
	}

	newNode := Node{
		value: value,
		next:  nil,
	}
	current := l.head
	for current != nil {
		current = current.next
	}

	current = &newNode
	return
}
