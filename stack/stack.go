package stack

import "errors"

type stack[T any] struct {
	internalSlice     []T
	internalSliceSize int
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{
		internalSlice:     make([]T, 0, 10),
		internalSliceSize: 0,
	}
}

func (s *stack[T]) Push(val T) {
	s.internalSlice = append(s.internalSlice, val)
	s.internalSliceSize++
}

func (s *stack[T]) Pop() error {
	if s.internalSliceSize == 0 {
		return errors.New("stack is empty")
	}
	s.internalSlice = s.internalSlice[:s.internalSliceSize]
	s.internalSliceSize--
	return nil
}

func (s *stack[T]) Top() (T, error) {
	if s.internalSliceSize != 0 {
		return s.internalSlice[s.internalSliceSize-1], nil
	} else {
		return *new(T), errors.New("stack is empty")
	}
}

func (s *stack[T]) IsEmpty() bool {
	return s.internalSliceSize == 0
}
