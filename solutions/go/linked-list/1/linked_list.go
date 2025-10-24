package linkedlist

import "errors"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.
type Node struct {
	Value interface{}
	prev *Node
	next *Node
}

type List struct {
	first *Node
	last *Node
}

func NewList(elements ...interface{}) *List {
	list := List{
		first: nil,
		last: nil,
	}

	if len(elements) == 0 {
		return &list
	}

	list.first = &Node{
		Value: elements[0],
		prev: nil,
		next: nil,
	}
	current := list.first
	
	for _, elem := range elements[1:] {
		newNode := &Node{
			Value: elem,
			prev: current,
			next: nil,
		}
		current.next = newNode
		current = newNode
	}

	list.last = current

	return &list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{
		Value: v,
		prev: nil,
		next: l.first,
	}

	if l.first == nil {
		l.first = newNode
		l.last = newNode
	} else {
		l.first.prev = newNode
		l.first = newNode
	}
}

func (l *List) Push(v interface{}) {
	newNode := &Node{
		Value: v,
		prev: l.last,
		next: nil,
	}
	if l.last == nil {
		l.first = newNode
		l.last = newNode
	} else {
		l.last.next = newNode
		l.last = newNode
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.first == nil {
		return nil, errors.New("cannot Shift an empty list")
	}
	if l.first.next == nil {
		v := l.first.Value
		l.first = nil
		l.last = nil
		return v, nil
	}

	first := l.first
	second := first.next

	l.first = second
	second.prev = nil

	first.next = nil

	return first.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.last == nil {
		return nil, errors.New("cannot Pop an empty list")
	}
	if l.last.prev == nil {
		v := l.first.Value
		l.first = nil
		l.last = nil
		return v, nil
	}
	last := l.last
	secondLast := last.prev

	l.last = secondLast
	secondLast.next = nil

	last.prev = nil
	return last.Value, nil
}

func (l *List) Reverse() {
	head := l.first
	current := head

	var tail *Node = nil
	for current != nil {
		next := current.next
		current.next = current.prev
		current.prev = next
		tail = current
		current = next	
	}

	l.last = head	
	l.first = tail
}

func (l *List) First() *Node {
	if l == nil {
		return nil
	}
	return l.first
}

func (l *List) Last() *Node {
	if l == nil {
		return nil
	}
	return l.last
}
