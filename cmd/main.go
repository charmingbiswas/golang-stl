// This package main will only be used for manual testing of implemented Golang data structures and algorithms
package main

import (
	"fmt"
	"reflect"
)

// Do manual testing of your data structures and algorithms here
func main() {
	v := []int{7}
	fmt.Println(reflect.DeepEqual(v, []int{7}))
}
