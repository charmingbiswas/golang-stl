package graph

// WIP - Work in progress

import (
	"cmp"
	"container/list"
)

type graphAdjList[T cmp.Ordered] struct {
	data map[T]list.List
}

type graphAdjMaxtrix[T cmp.Ordered] struct {
	data [][]T
}

func NewAdjListGraph[T cmp.Ordered]() *graphAdjList[T] {
	return &graphAdjList[T]{data: make(map[T]list.List)}
}

func NewAdjMatrixGraph[T cmp.Ordered]() *graphAdjMaxtrix[T] {
	return &graphAdjMaxtrix[T]{data: make([][]T, 0)}
}

func (g *graphAdjList[T]) AddVertex(value T) {
	g.data[value] = *list.New()
}

func (g *graphAdjMaxtrix[T]) AddVertex(value T) {

	g.data = append(g.data)
}
