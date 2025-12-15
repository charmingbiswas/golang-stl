package stack

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, 10),
	}
}

func (st *Stack[T]) Push(val T) {
	st.data = append(st.data, val)
}

func (st *Stack[T]) Pop() {
	if len(st.data) == 0 {
		return
	}
	st.data = st.data[:len(st.data)-1]
}

func (st *Stack[T]) Top() T {
	if len(st.data) == 0 {
		return *new(T)
	}

	return st.data[len(st.data)-1]
}

func (st *Stack[T]) IsEmpty() bool {
	return len(st.data) == 0
}

func (st *Stack[T]) Size() int {
	return len(st.data)
}
