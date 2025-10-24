package linkedlist

import "errors"

// Define the List and Element types here.
type Element struct {
	ID int
	next *Element
}

type List struct {
	head *Element
}

func New(elements []int) *List {
	list := &List{}
	if len(elements) == 0 {
		return list
	}

	current := &Element{
		ID: elements[0],
		next: nil,
	}
	list.head = current

	for _, val := range elements[1:] {
		current.next = &Element{
			ID: val,
			next: nil,
		}
		current = current.next
	}

	return list
}

func (l *List) Size() int {
	if l == nil {
		return 0
	}

	size := 0
	current := l.head

	for current != nil {
		size++
		current = current.next
	}

	return size
}

func (l *List) Push(element int) {
	node := &Element{
		ID: element,
		next: nil,
	}

	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

func (l *List) Pop() (int, error) {
	// no items
	if l.head == nil {
		return 0, errors.New("empty list")
	}
	// one item
	if l.head.next == nil {
		v := l.head.ID
		l.head = nil

		return v, nil
	}

	// multiple items
	prev := l.head
	current := prev.next
	
	for current.next != nil {
		prev = current
		current = current.next
	}

	prev.next = nil
	
	return current.ID, nil
}

func (l *List) Array() []int {
	if l == nil {
		return []int{}
	}

	arr := make([]int, l.Size())
	
	current := l.head
	i := 0
	
	for current != nil {
		arr[i] = current.ID
		current = current.next
		i++
	}

	return arr
}

func (l *List) Reverse() *List {
	// empty list
	if l.head == nil {
		return nil
	}
	// one item 
	if l.head.next == nil {
		return l
	}

	prev := l.head
	current := l.head.next
	prev.next = nil

	for current != nil {
		next := current.next 
		current.next = prev 
		prev = current
		current = next
	}

	l.head = prev
	return l
}
