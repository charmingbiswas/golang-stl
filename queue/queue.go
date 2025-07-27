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

func (this *queue[T]) PopFront() {
	if len(this.data) == 0 {
		return
	}
	this.data = this.data[1:]
}

func (this *queue[T]) PopBack() {
	if len(this.data) == 0 {
		return
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *queue[T]) Front() T {
	if len(this.data) == 0 {
		return *new(T)
	}

	return this.data[0]
}

func (this *queue[T]) Back() T {
	if len(this.data) == 0 {
		return *new(T)
	}

	return this.data[len(this.data)-1]
}

func (this *queue[T]) IsEmpty() bool {
	return len(this.data) == 0
}
