// This package contains implementation of most popular tree data structures
package trees

// For better understanding of theory and algorithm, please refer to this playlist:
// https://www.youtube.com/playlist?list=PL9xmBV_5YoZNqDI8qfOZgzbqahCUmUEin

import (
	"cmp"
)

type color string

const (
	RED   color = "red"
	BLACK color = "black"
)

type node[T cmp.Ordered, V any] struct {
	Key       T
	Val       V
	NodeColor color
	Left      *node[T, V]
	Right     *node[T, V]
	Parent    *node[T, V]
}

type rbTree[T cmp.Ordered, V any] struct {
	Root *node[T, V]
	NIL  *node[T, V]
}

func NewRedBlackTree[T cmp.Ordered, V any]() *rbTree[T, V] {
	nilNode := &node[T, V]{
		NodeColor: BLACK,
	}
	return &rbTree[T, V]{
		Root: nilNode,
		NIL:  nilNode,
	}
}

func (t *rbTree[T, V]) Insert(key T, value V) {

	// create a new node
	newNode := &node[T, V]{
		Key:       key,
		Val:       value,
		NodeColor: RED,
		Left:      t.NIL,
		Right:     t.NIL,
		Parent:    t.NIL,
	}

	parent := t.NIL

	// start from root and keep traversing the tree until you find the correct spot for insertion
	currentNode := t.Root

	for currentNode != t.NIL {
		parent = currentNode
		if newNode.Key < currentNode.Key {
			currentNode = currentNode.Left
		} else if newNode.Key > currentNode.Key {
			currentNode = currentNode.Right
		} else if newNode.Key == currentNode.Key {
			currentNode.Val = value
			return
		}
	}

	newNode.Parent = parent

	if parent == t.NIL {
		t.Root = newNode
	} else if newNode.Key < parent.Key {
		parent.Left = newNode
	} else if newNode.Key > parent.Key {
		parent.Right = newNode
	}

	t.insertFixup(newNode)
}

func (t *rbTree[T, V]) Delete(key T) bool {
	node := t.search(key)

	if node == t.NIL {
		return false
	}

	t.deleteNode(node)
	return true
}

func (t *rbTree[T, V]) Search(key T) (V, bool) {
	n := t.search(key)
	if n == t.NIL {
		return t.NIL.Val, false
	}

	return n.Val, true
}

func (t *rbTree[T, V]) IsEmpty() bool {
	return t.Root == t.NIL
}

func (t *rbTree[T, V]) PrintInOrder() {}

func (t *rbTree[T, V]) insertFixup(n *node[T, V]) {
	// check if inserted node's parent color is RED
	// meaning newly inserted node is NOT the root node
	for n.Parent.NodeColor == RED {
		// if parent of inserted node is left child of grandparent
		if n.Parent == n.Parent.Parent.Left {
			uncle := n.Parent.Parent.Right
			if uncle.NodeColor == RED {
				// Case 1: Uncle is red
				// Change uncle to black, parent to black and grandparent to red
				n.Parent.NodeColor = BLACK
				uncle.NodeColor = BLACK
				n.Parent.Parent.NodeColor = RED
				n = n.Parent.Parent
			} else {
				if n == n.Parent.Right {
					// Case 2: Uncle is black and node is right child
					// Meaning: Triangle is being formed because node is right child but node's parent is left child
					// Rotate in OPPOSITE direction of node at the parent node
					// In this case, since node is right child, rotate left at the parent
					/*
							Only for understanding and reference, not ideal example
							 5
							/ \
						   3   9
							\
							 4
					*/
					n = n.Parent
					t.leftRotate(n)
				}
				// Case 3: Uncle is black and node is left child
				// Meaning: Straing line is being formed because node is left child and parent is also left child
				// Rotate in opposite direction of node at the parent node
				// In this case, since node is left child, rotate right at the parent
				// And re-color parent to Black and grandparent to Red to preserve Red Black Tree property
				// Re-coloring needs to be done in both cases, hence added outside the if condition
				/*
						Only for understanding and reference, not ideal example
						  5
						 / \
					    3   9
					   /
					  4
				*/
				n.Parent.NodeColor = BLACK
				n.Parent.Parent.NodeColor = RED
				t.rightRotate(n.Parent.Parent)
			}
		} else { // if parent of inserted node is right child of grandparent
			uncle := n.Parent.Parent.Left
			if uncle.NodeColor == RED {
				// Case 1: Uncle is red
				n.Parent.NodeColor = BLACK
				uncle.NodeColor = BLACK
				n.Parent.Parent.NodeColor = RED
				n = n.Parent.Parent
			} else {
				if n == n.Parent.Left {
					// Case 2: Uncle is black and node is left child
					n = n.Parent
					t.rightRotate(n)
				}
				// Case 3: Uncle is black and node is right child
				n.Parent.NodeColor = BLACK
				n.Parent.Parent.NodeColor = RED
				t.leftRotate(n.Parent.Parent)
			}
		}
	}
	t.Root.NodeColor = BLACK // ROOT is always black
}

func (t *rbTree[T, V]) deleteNode(n *node[T, V]) {

	originalNode := n
	originalColor := originalNode.NodeColor

	var replacementNode *node[T, V]

	// There are 3 scenarios which we need to consider
	// If left child of node to be deleted is NIL
	// If right child of node to be deleted is NIL
	// If neither is the case and both children exist

	if n.Left == t.NIL {
		// Case 1: NIL left child
		// Transplant node n with it's RIGHT child
		replacementNode = n.Right
		t.transplant(n, n.Right)
	} else if n.Right == t.NIL {
		// Case 2: NIL right child
		// Transplant node n with it's LEFT child
		replacementNode = n.Left
		t.transplant(n, n.Left)
	} else {
		// Case 3: Neither children are NIL
		// Now since both chilren exist, we need to first find the minimum in the RIGHT sub tree

		successorNode := t.minimum(n.Right)
		originalNode = successorNode
		originalColor = successorNode.NodeColor
		replacementNode = successorNode.Right

		if successorNode.Parent == n {
			replacementNode.Parent = successorNode
		} else {
			t.transplant(n, successorNode)
			successorNode.Right = n.Right
			successorNode.Right.Parent = successorNode
		}

		t.transplant(n, successorNode)
		successorNode.Left = n.Left
		successorNode.Left.Parent = successorNode
		successorNode.NodeColor = n.NodeColor

	}

	if originalColor == BLACK {
		t.deleteFixup(replacementNode)
	}
}

func (t *rbTree[T, V]) deleteFixup(n *node[T, V]) {
	/*
		There are 4 types of fixes that we will encounter:
		We will call sibling of node n as m
		1. When m is RED
		2. When m is BLACK and both it's children are BLACK
		3. When m is BLACK and it's right child is BLACK but left child is RED
		4. When m is BLACK and it's right is RED but left child is BLACK
		// There conditions are not mutually exclusive and we need to call multiple fix ups to maintain RED BLACK TREE behavior
	*/

	// We use a for loop since there can be multiple fix ups needed as long as RED BLACK TREE condition is not satified

	for n != t.NIL && n.NodeColor == BLACK {
		if n == n.Parent.Left {
			sibling := n.Parent.Right // sibling will be right child since node n is left child

			// CASE 1
			if sibling.NodeColor == RED {
				// 1. Color sibling BLACK
				// 2. Color parent RED
				// 3. LEFT rotate on parent
				// 4. Change pointer back to new sibling of node n

				sibling.NodeColor = BLACK
				n.Parent.NodeColor = RED
				t.leftRotate(n.Parent)
				sibling = n.Parent.Right
			}

			// CASE 2
			if sibling.Left.NodeColor == BLACK && sibling.Right.NodeColor == BLACK {
				// 1. Change sibling color to RED
				// 2. Move pointer to parent
				sibling.NodeColor = RED
				n = n.Parent
			} else {
				// CASE 3
				if sibling.Right.NodeColor == BLACK {
					// 1. Change left child of sibling to black
					// 2. Change sibling to RED
					// 3. Do RIGHT rotation since sibling is right child
					// 4. Change pointer back to new sibling after the rotation

					sibling.Left.NodeColor = BLACK
					sibling.NodeColor = RED
					t.rightRotate(sibling)
					sibling = n.Parent.Right
				}

				// CASE 4
				// 1. Set sibling color to color of node n's parent
				// 2. Set node n's parent color to BLACK
				// 3. Set right child of sibling to BLACK
				// 4. Left rotate on node n's parent
				// 5. Set node pointer to root of the tree

				sibling.NodeColor = n.Parent.NodeColor
				n.Parent.NodeColor = BLACK
				sibling.Right.NodeColor = BLACK
				t.leftRotate(n.Parent)
				n = t.Root
			}
		} else {
			// same logic, just reverse the directions of rotations
			sibling := n.Parent.Left // sibling will be right child since node n is left child

			// CASE 1
			if sibling.NodeColor == RED {

				sibling.NodeColor = BLACK
				n.Parent.NodeColor = RED
				t.rightRotate(n.Parent)
				sibling = n.Parent.Left
			}

			// CASE 2
			if sibling.Right.NodeColor == BLACK && sibling.Left.NodeColor == BLACK {
				sibling.NodeColor = RED
				n = n.Parent
			} else {
				// CASE 3
				if sibling.Left.NodeColor == BLACK {

					sibling.Right.NodeColor = BLACK
					sibling.NodeColor = RED
					t.leftRotate(sibling)
					sibling = n.Parent.Left
				}

				// CASE 4
				sibling.NodeColor = n.Parent.NodeColor
				n.Parent.NodeColor = BLACK
				sibling.Left.NodeColor = BLACK
				t.rightRotate(n.Parent)
				n = t.Root
			}
		}
	}

	n.NodeColor = BLACK // final step
}

func (t *rbTree[T, V]) leftRotate(n *node[T, V]) {
	/*
			  1
			 / \
			3   4
		       / \
			  2   5

			We are rotating around 4.
			After rotate left, the structure would be like so

			    4
			   /  \
			  1    5
			 / \
			3   2
	*/

	x := n       // 1
	y := n.Right // 4

	x.Right = y.Left // NIL in this case

	if y.Left != t.NIL {
		y.Left.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == t.NIL {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}

	y.Left = x
	x.Parent = y
}

func (t *rbTree[T, V]) rightRotate(n *node[T, V]) {
	// Same logic as left rorate, just flip left child with right child and vice versa
	x := n
	y := n.Left

	x.Left = y.Right // NIL in this case

	if y.Right != t.NIL {
		y.Right.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == t.NIL {
		t.Root = y
	} else if x == x.Parent.Right {
		x.Parent.Right = y
	} else {
		x.Parent.Left = y
	}

	y.Right = x
	x.Parent = y
}

func (t *rbTree[T, V]) search(key T) *node[T, V] {
	current := t.Root
	for current != t.NIL {
		if key == current.Key {
			return current
		} else if key < current.Key {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return t.NIL
}

func (t *rbTree[T, V]) transplant(n, m *node[T, V]) {
	// We need to deal with three cases
	// If n is root node
	// If n is left child of parent
	// If n is right child of parent

	if n.Parent == t.NIL {
		// meaning n is root node, there is no parent
		// just make m as the new root
		t.Root = m
	} else if n == n.Parent.Left {
		// n is left child
		// transplat with m
		n.Parent.Left = m
	} else {
		n.Parent.Right = m
	}

	n.Parent = m.Parent
}

// Returns the minimum key in a sub tree rooted at node n
func (t *rbTree[T, V]) minimum(n *node[T, V]) *node[T, V] {
	for n.Left != t.NIL {
		n = n.Left
	}
	return n
}
