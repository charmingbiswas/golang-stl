package queue

type queue[T any] struct {
	data []T
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{
		data: make([]T, 0, 10),
	}
}

func (this *queue[T]) PushFront(val T) {
	this.data = append([]T{val}, this.data...)
}

func (this *queue[T]) PushBack(val T) {
	this.data = append(this.data, val)
}

func (this *queue[T]) PopFront() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[1:]
	return true
}

func (this *queue[T]) PopBack() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[:len(this.data)-1]
	return true
}

func (this *queue[T]) Front() (T, bool) {
	if len(this.data) == 0 {
		return *new(T), false
	}

	return this.data[0], true
}

func (this *queue[T]) Back() (T, bool) {
	if len(this.data) == 0 {
		return *new(T), false
	}

	return this.data[len(this.data)-1], true
}

func (this *queue[T]) IsEmpty() bool {
	return len(this.data) == 0
}
