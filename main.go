// This package is for quick manual test of your data structures
package main

import (
	"fmt"

	"github.com/charmingbiswas/golang-stl/algo"
)

func main() {
	v := []int{-1, -5}
	fmt.Println(algo.LowerBound(v, 1))
}
