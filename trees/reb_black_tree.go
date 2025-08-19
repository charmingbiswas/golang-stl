// This package contains implementation of most popular tree data structures
package trees

import "cmp"

type color string

const (
	RED   color = "red"
	BLACK color = "black"
)

type node[T cmp.Ordered, V any] struct {
	key       T
	val       V
	nodeColor color
	left      *node[T, V]
	right     *node[T, V]
	parent    *node[T, V]
}

type rbTree[T cmp.Ordered, V any] struct {
	root *node[T, V]
	NIL  *node[T, V]
}

func NewRedBlackTree[T cmp.Ordered, V any]() *rbTree[T, V] {
	nilNode := &node[T, V]{
		nodeColor: BLACK,
	}
	return &rbTree[T, V]{
		root: nilNode,
		NIL:  nilNode,
	}
}

func (this *rbTree[T, V]) Insert(key T, value V) {
	// initialize new node
	newNode := &node[T, V]{
		key:       key,
		val:       value,
		nodeColor: RED,
		left:      this.NIL,
		right:     this.NIL,
		parent:    this.NIL,
	}

	var parent *node[T, V] = this.NIL
	current := this.root

	// standard BST insertion
	for current != this.NIL {
		parent = current
		if newNode.key < current.key {
			current = current.left
		} else if newNode.key > current.key {
			current = current.right
		} else {
			current.val = value
		}
	}

	newNode.parent = parent

	if parent == this.NIL {
		this.root = newNode
	} else if newNode.key < parent.key {
		parent.left = newNode
	} else if newNode.key > parent.key {
		parent.right = newNode
	}

	// check for violations for red black tree and fix them
	this.insertFixup(newNode)
}

func (this *rbTree[T, V]) insertFixup(node *node[T, V]) {
	// 4 scenarios
	// 1. Z = root -> change color to black
	// 2. Z.uncle = red -> recolor parent, grandparent and uncle
	// 3. Z.uncle = black (line) -> rotate Z.parent opposite direction of Z
	// 4. Z.uncle = black (triangle) -> rotate Z.grandparent in opposite direction of z and recolor

}
