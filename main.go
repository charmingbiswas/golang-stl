// This package is for quick manual test of your data structures
package main

import (
	"container/list"
)

type LRU struct {
	data     map[int]int
	capacity int
	priority list.List
}

func NewCache(capacity int) *LRU {
	return &LRU{
		data:     make(map[int]int, 0),
		capacity: capacity,
	}
}

// func (this *LRU) Get(key int) int {

// }

// func (this *LRU) Set(key int, val int) {
// }

func main() {
}
