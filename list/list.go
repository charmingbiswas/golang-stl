// This package includes implementation of doubly linked list
package list

// Underlying node struct type that forms the list
type Node[T any] struct {
	val  T
	next *Node[T]
	prev *Node[T]
}

type list[T any] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{
		val:  val,
		next: nil,
		prev: nil,
	}
}

func NewList[T any]() *list[T] {
	return &list[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (this *list[T]) PushFront(val T) {
	node := NewNode(val)
	if this.head == nil {
		this.head = node
		this.tail = node
		this.size++
	} else {
		prev := this.head
		this.head = node
		this.head.next = prev
		this.size++
	}
}

func (this *list[T]) PushBack(val T) {
	node := NewNode(val)
	if this.head == nil {
		this.head = node
		this.tail = node
		this.size++
	} else {
		prev := this.tail
		prev.next = node
		this.tail = node
		this.tail.prev = prev
		this.size++
	}
}

func (this *list[T]) Front() (T, bool) {
	if this.size == 0 {
		return *new(T), false
	}

	return this.head.val, true
}

func (this *list[T]) Back() (T, bool) {
	if this.size == 0 {
		return *new(T), false
	}

	return this.tail.val, true
}

func (this *list[T]) IsEmpty() bool {
	return this.size == 0
}
