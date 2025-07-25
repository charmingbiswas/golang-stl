// This package main will only be used for manual testing of implemented Golang data structures and algorithms
package main

import (
	"fmt"

	"github.com/charmingbiswas/golang-stl/heap"
)

type Pair struct {
	first  int
	second string
}

func main() {
	maxHeap := heap.NewHeapWithFunc(func(a, b Pair) bool { return a.first > b.first })

	maxHeap.Push(Pair{first: 1, second: "one"})
	maxHeap.Push(Pair{first: 2, second: "two"})
	maxHeap.Push(Pair{first: 3, second: "three"})
	maxHeap.Push(Pair{first: 4, second: "four"})
	maxHeap.Push(Pair{first: 5, second: "five"})

	fmt.Println(maxHeap)
	maxHeap.Pop()
	fmt.Println(maxHeap)

}
