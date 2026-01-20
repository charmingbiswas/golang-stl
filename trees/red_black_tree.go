// This package contains implementation of most popular tree data structures
package trees

// For better understanding of theory and algorithm, please refer to this playlist:
// https://www.youtube.com/playlist?list=PL9xmBV_5YoZNqDI8qfOZgzbqahCUmUEin

/*
	1. Red black tree is a self balancing binary search tree.
	2. The root node must be colored black.
	3. The children of red color node must be colored black. There should not be two consecutive red nodes.
	4. In all the paths of the tree there should be same number of black color nodes.
	5. Every new node must be inserted with red color.
	6. Every leaf( i.e nil node) must be colored black.
*/

import (
	"cmp"
	"fmt"
	"iter"
)

type Color string

const (
	RED   Color = "red"
	BLACK Color = "black"
)

type RedBlackTreeNode[T cmp.Ordered, V any] struct {
	Key       T
	Value     V
	NodeColor Color
	Left      *RedBlackTreeNode[T, V]
	Right     *RedBlackTreeNode[T, V]
	Parent    *RedBlackTreeNode[T, V]
}

type RedBlackTree[T cmp.Ordered, V any] struct {
	Root     *RedBlackTreeNode[T, V]
	NIL      *RedBlackTreeNode[T, V]
	treeSize int
}

// Returns a pointer to an instance of a RedBlackTree struct.
func NewRedBlackTree[T cmp.Ordered, V any]() *RedBlackTree[T, V] {
	nilNode := &RedBlackTreeNode[T, V]{
		NodeColor: BLACK,
	}

	return &RedBlackTree[T, V]{
		Root:     nilNode,
		NIL:      nilNode,
		treeSize: 0,
	}
}

// Inserts a key-value pair into the RedBlackTree.
// The key has to be of type cmp.Ordered, value can be anything.
func (t *RedBlackTree[T, V]) Insert(key T, value V) {
	newNode := &RedBlackTreeNode[T, V]{
		NodeColor: RED,
		Key:       key,
		Value:     value,
		Left:      t.NIL,
		Right:     t.NIL,
		Parent:    t.NIL,
	}

	currentNode := t.Root
	parentNode := t.NIL

	// Traverse tree
	for currentNode != t.NIL {
		/*
			NIL (parentNode)
			 |
			Root (currentNode)
		*/
		parentNode = currentNode

		/*
			NIL
			 |
			Root (parent)
			 |
			(currentNode) Traverse tree using this node
		*/
		if newNode.Key < currentNode.Key {
			currentNode = currentNode.Left
		} else if newNode.Key > currentNode.Key {
			currentNode = currentNode.Right
		} else {
			// if exact key is found
			// update the value and return, nothing left to do
			currentNode.Value = value
			return
		}
	}

	// now need to position the parent node pointer to correct parent
	// this case will happen when exact key is not found
	// so need to insert the newNode somewhere at the correct position
	newNode.Parent = parentNode // this case, parentNode is t.Nil if tree is empty

	if parentNode == t.NIL {
		// is parent is nil, insert new node at root
		t.Root = newNode
	} else if newNode.Key < parentNode.Key {
		parentNode.Left = newNode
	} else if newNode.Key > parentNode.Key {
		parentNode.Right = newNode
	}

	// after insertion, call insert fixup helper to maintain red black tree properties
	t.insertFixup(newNode)
	t.treeSize++
}

// Deletes a key from the tree.
// Key has to be of type cmp.Ordered
// If key does not exist, returns false, otherwise true if deletion is successful.
func (t *RedBlackTree[T, V]) Delete(key T) bool {
	/*
		For deletion we need to consider these cases
		1. Left child of node to be deleted is NIL
		2. Right child of node to be delete is NIL
		3. Neither are NIL

		IMPORTANT:
		If we deleted a node with original color as RED, NO FIXUP is needed.
		Since deleting a red node does not reduce the number of black nodes on any path
		Fixup is called ONLY for deleting node which have BLACK color
	*/

	nodeToBeDeleted, ok := t.Search(key)
	if !ok {
		return false // node with given key does not exist on the tree
	}

	originalNode := nodeToBeDeleted
	originalNodeColor := nodeToBeDeleted.NodeColor
	var replacementNode *RedBlackTreeNode[T, V]

	if originalNode.Left == t.NIL {
		// if left child is NIL, transplant this node with it's right child
		replacementNode = originalNode.Right
		t.transplant(originalNode, originalNode.Right)
	} else if originalNode.Right == t.NIL {
		replacementNode = originalNode.Left
		t.transplant(originalNode, originalNode.Left)
	} else {
		// if neither children are NIL
		// find inorder successor, i.e the smallest node in the right subtree
		successor := t.minimum(originalNode.Right)
		originalNode = successor
		originalNodeColor = successor.NodeColor
		replacementNode = successor.Right

		/*
			Tree would be looking like this (we are trying to delete 12)
				12
			   /   \
			  8    15
			 / \   / \
			1  9  13 23
			    \  \
				10 NIL

			successor node is 13, replacements node becomes successor.Right = t.NIL
			After transplant, this is the new tree

			    12  (13)(successor, floating around)
			   /   \
			  8    15
			 / \   / \
			1  9  NIL  23
			    \
				10
		*/

		if successor.Parent == nodeToBeDeleted {
			replacementNode.Parent = successor
		} else {
			t.transplant(successor, successor.Right) // transplant successor with it's right child, which would be NIL child
			successor.Right = nodeToBeDeleted.Right
			successor.Right.Parent = successor
		}

		// now transplant successor with nodeToBeDeleted and change the pointers around
		t.transplant(nodeToBeDeleted, successor)
		successor.Left = nodeToBeDeleted.Left
		successor.Left.Parent = successor
		successor.NodeColor = nodeToBeDeleted.NodeColor // keep the color same
	}

	if originalNodeColor == BLACK {
		// fixup is only needed when deleting a black node, if you delete a red node the number of black nodes per path does not change
		// hence no fixup needed for RED node deletions
		t.deleteFixup(replacementNode)
	}
	t.treeSize--
	return true
}

// Searches for a key in the tree.
// Key has to be of type cmp.Ordered
// Returns the node and boolean value.
// Boolean is true if key is found, otherwise false.
func (t *RedBlackTree[T, V]) Search(key T) (*RedBlackTreeNode[T, V], bool) {
	currentNode := t.Root

	for currentNode != t.NIL {
		if currentNode.Key == key {
			return currentNode, true
		} else if key < currentNode.Key {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}

	return t.NIL, false
}

// Returns true if tree is empty, otherwise false.
func (t *RedBlackTree[T, V]) IsEmpty() bool {
	return t.Root == t.NIL
}

// Prints the key value pairs in the tree.
// The ordering is 'InOrder' sorted ordering.
func (t *RedBlackTree[T, V]) PrintTree() {
	// Using Morris Traversal to preserve memory space
	currentNode := t.Root
	for currentNode != t.NIL {
		// check if left child exists
		if currentNode.Left == t.NIL {
			// No left child, print current node value.
			// Move on to right child.
			fmt.Printf("KEY: %v, VALUE: %v, NODE COLOR: %v\n", currentNode.Key, currentNode.Value, currentNode.NodeColor)
			currentNode = currentNode.Right
		} else {
			// left child exists
			// find inorder predecessor
			// meaning right most child of left child
			predecessor := currentNode.Left
			for predecessor.Right != t.NIL && predecessor.Right != currentNode {
				predecessor = predecessor.Right
			}

			switch predecessor.Right {
			case t.NIL:
				predecessor.Right = currentNode
				currentNode = currentNode.Left
			case currentNode:
				predecessor.Right = t.NIL
				fmt.Printf("KEY: %v, VALUE: %v, NODE COLOR: %v\n", currentNode.Key, currentNode.Value, currentNode.NodeColor)
				currentNode = currentNode.Right
			}
		}
	}
}

// Returns a 'push' iterator to the tree.
// Works with 'for range' expression.
// Returns key-value pair per iteration in 'InOrder' sorted ordering.
// Values are returned starting from the first index.
func (t *RedBlackTree[T, V]) ForwardIterator() iter.Seq2[T, V] {
	type Pair struct {
		key   T
		value V
	}

	result := make([]Pair, 0, t.treeSize)

	// Use Morris Traversal for space efficient traversal
	currentNode := t.Root
	for currentNode != t.NIL {
		if currentNode.Left == t.NIL {
			result = append(result, Pair{key: currentNode.Key, value: currentNode.Value})
			currentNode = currentNode.Right
		} else {
			inOrderPredecessor := currentNode.Left
			for inOrderPredecessor.Right != t.NIL && inOrderPredecessor.Right != currentNode {
				inOrderPredecessor = inOrderPredecessor.Right
			}

			switch inOrderPredecessor.Right {
			case t.NIL:
				inOrderPredecessor.Right = currentNode
				currentNode = currentNode.Left
			case currentNode:
				inOrderPredecessor.Right = t.NIL
				result = append(result, Pair{key: currentNode.Key, value: currentNode.Value})
				currentNode = currentNode.Right
			}
		}
	}

	return func(yield func(T, V) bool) {
		for _, pair := range result {
			if !yield(pair.key, pair.value) {
				return
			}
		}
	}
}

// Returns a 'push' iterator to the tree.
// Works with 'for range' expression.
// Returns key-value pair per iteration in 'InOrder' sorted ordering.
// Values are returned from the starting from the last index.
func (t *RedBlackTree[T, V]) BackwardIterator() iter.Seq2[T, V] {
	type Pair struct {
		key   T
		value V
	}

	result := make([]Pair, 0, t.treeSize)

	// Use Morris Traversal for space efficient traversal
	currentNode := t.Root
	for currentNode != t.NIL {
		if currentNode.Left == t.NIL {
			result = append(result, Pair{key: currentNode.Key, value: currentNode.Value})
			currentNode = currentNode.Right
		} else {
			inOrderPredecessor := currentNode.Left
			for inOrderPredecessor.Right != t.NIL && inOrderPredecessor.Right != currentNode {
				inOrderPredecessor = inOrderPredecessor.Right
			}

			switch inOrderPredecessor.Right {
			case t.NIL:
				inOrderPredecessor.Right = currentNode
				currentNode = currentNode.Left
			case currentNode:
				inOrderPredecessor.Right = t.NIL
				result = append(result, Pair{key: currentNode.Key, value: currentNode.Value})
				currentNode = currentNode.Right
			}
		}
	}

	return func(yield func(T, V) bool) {
		for i := len(result) - 1; i >= 0; i-- {
			if !yield(result[i].key, result[i].value) {
				return
			}
		}
	}
}

// Returns the current number of nodes in the tree.
func (t *RedBlackTree[T, V]) Size() int {
	return t.treeSize
}

// Clears and resets the tree to an empty tree.
func (t *RedBlackTree[T, V]) Clear() {
	t.Root = t.NIL
}

func (t *RedBlackTree[T, V]) insertFixup(node *RedBlackTreeNode[T, V]) {
	// 1. Check if newly inserted node's parent color is RED
	// if not, then new node is the root node
	for node.Parent.NodeColor == RED {
		// if parent of inserted node is left child of it's grandparent
		/*
					ROOT
					|
				GRANDPARENT
				/    	   \
			  PARENT	   UNCLE
			  /
			INSERTED NODE
		*/
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			/*
				1. Case One: uncle is RED
				-> Change color of uncle and parent to BLACK and grandparent to RED
				-> Move pointer from current node to grandparent of current node
			*/
			if uncle.NodeColor == RED {
				uncle.NodeColor = BLACK
				node.Parent.NodeColor = BLACK
				node.Parent.Parent.NodeColor = RED
				node = node.Parent.Parent
			} else {
				/*
					2. Case Two: uncle is black and current node is RIGHT child of parent // triangle is formed between inserted node, parent and grandparent
								ROOT
								|
							GRANDPARENT
					(triangle)/    	   \
						PARENT	      UNCLE
							\
							INSERTED NODE
					-> Rotate on parent in the OPPOSITE direction of the current node.
					-> change color of parent to BLACK and grand parent to RED
					-> color change is needed in both case 2 and 3, hence done outside the if condition
				*/
				if node == node.Parent.Right {
					node = node.Parent // this is done, because after rotation, the parent will become the leaf node
					t.rotateLeft(node) // rorate on the parent for triangle
				}

				/*
						3. Case Three: uncle is black and current node is RIGHT child of parent // straight line is formed
									  ROOT
										|
									GRANDPARENT
					(straigth line)/    	   \
								PARENT	      UNCLE
								/
								INSERTED NODE
						-> rotate in opposite direction on the 'grandparent' node
						-> change color of parent to black and grandparent to red
				*/
				node.Parent.NodeColor = BLACK
				node.Parent.Parent.NodeColor = RED
				t.rotateRight(node.Parent.Parent) // rotate on the grandparent for straight line
			}
		} else {
			// inserted node's parent if the RIGHT child of grandparent
			/*
						ROOT
						|
					GRANDPARENT
					/    	   \
				  UNCLE	      PARENT
				  				\
							INSERTED NODE
			*/
			// Again check the same conditions as before
			uncle := node.Parent.Parent.Left

			if uncle.NodeColor == RED {
				uncle.NodeColor = BLACK
				node.Parent.NodeColor = BLACK
				node.Parent.Parent.NodeColor = RED
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					t.rotateRight(node)
				}

				node.Parent.NodeColor = BLACK
				node.Parent.Parent.NodeColor = RED
				t.rotateLeft(node.Parent.Parent)
			}
		}
	}

	t.Root.NodeColor = BLACK // root is always black
}

func (t *RedBlackTree[T, V]) deleteFixup(node *RedBlackTreeNode[T, V]) {
	/*
		Fixes red black tree violations after deletion has taken place
		We need to understand certain cases
		sibling = sibling of the node on which delete fixup is being performed

		Case 1: sibling is RED
		Case 2: sibling is BLACK and both it's children are BLACK
		Case 3: sibling is BLACK and left child is RED but right child is BLACK
		Case 4: sibling is BLACK and right child is RED

		These cases are not exclusive and there can be multiple violations being fixed in the same function call
	*/

	for node != t.Root && node.NodeColor == BLACK {
		// if fixup node is the left child of it's parent
		switch node {
		case node.Parent.Left:
			sibling := node.Parent.Right
			// Case 1: sibling is RED
			if sibling.NodeColor == RED {
				sibling.NodeColor = BLACK
				node.Parent.NodeColor = RED
				t.rotateLeft(node.Parent)
				sibling = node.Parent.Right // sibling will change after rotation, set it back to node's sibling after rotation
			}

			// Case 2: sibling is BLACK
			if sibling.Left.NodeColor == BLACK && sibling.Right.NodeColor == BLACK {
				sibling.NodeColor = RED
				node = node.Parent // move current pointer from node, to node's parent
			} else {
				// Case 3: sibling right child is black and left is red
				if sibling.Right.NodeColor == BLACK {
					sibling.Left.NodeColor = BLACK
					sibling.NodeColor = RED
					t.rotateRight(sibling)
					sibling = node.Parent.Right // reset sibling pointer after rotation to correct position
				}

				// Case 4: sibling right child is RED
				sibling.NodeColor = node.Parent.NodeColor
				node.Parent.NodeColor = BLACK
				sibling.Right.NodeColor = BLACK
				t.rotateLeft(node.Parent)
				node = t.Root
			}

		case node.Parent.Right:
			// fixup node is the right child of parent
			sibling := node.Parent.Left
			// Case 1:
			if sibling.NodeColor == RED {
				sibling.NodeColor = BLACK
				node.Parent.NodeColor = RED
				t.rotateRight(node.Parent)
				sibling = node.Parent.Left
			}

			// Case 2:
			if sibling.Left.NodeColor == BLACK && sibling.Right.NodeColor == BLACK {
				sibling.NodeColor = RED
				node = node.Parent
			} else {
				// Case 3:
				if sibling.Left.NodeColor == BLACK {
					sibling.Right.NodeColor = BLACK
					sibling.NodeColor = RED
					t.rotateLeft(sibling)
					sibling = node.Parent.Left
				}

				// Case 4:
				sibling.NodeColor = node.Parent.NodeColor
				node.Parent.NodeColor = BLACK
				sibling.Left.NodeColor = BLACK
				t.rotateRight(node.Parent)
				node = t.Root
			}
		}
	}

	node.NodeColor = BLACK
}

func (t *RedBlackTree[T, V]) rotateLeft(node *RedBlackTreeNode[T, V]) {
	/*
				  ROOT
					|
				GRANDPARENT
		(triangle)/    	   \
			PARENT	      UNCLE
				\
				INSERTED NODE(x)
				            \
							(y)
	*/
	x := node
	y := node.Right

	// Put y's left subtree into x's right subtree
	x.Right = y.Left

	if y.Left != t.NIL {
		y.Left.Parent = x
	}
	y.Parent = x.Parent

	if x.Parent == t.NIL {
		// meaning x is root node
		// so make y the new root
		t.Root = y
	} else {
		switch x {
		case x.Parent.Left:
			// x was the left child
			// put y there
			x.Parent.Left = y
		case x.Parent.Right:
			x.Parent.Right = y
		}
	}

	// put x on y's left
	y.Left = x
	x.Parent = y
}

func (t *RedBlackTree[T, V]) rotateRight(node *RedBlackTreeNode[T, V]) {
	/*
				  ROOT
					|
				GRANDPARENT
		(straight)/    	   \
			PARENT	      UNCLE
			/
		INSERTED NODE(x)
		  /
		(y)

	*/
	x := node
	y := node.Left

	x.Left = y.Right

	if y.Right != t.NIL {
		y.Right.Parent = x
	}

	y.Parent = x.Parent

	if x.Parent == t.NIL {
		t.Root = y
	} else {
		switch x {
		case x.Parent.Left:
			x.Parent.Left = y
		case x.Parent.Right:
			x.Parent.Right = y
		}
	}

	y.Right = x
	x.Parent = y
}

func (t *RedBlackTree[T, V]) transplant(n, m *RedBlackTreeNode[T, V]) {
	// We need to deal with three cases
	// If n is root node
	// If n is left child of parent
	// If n is right child of parent

	// 1. n is root node
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

	// take n's parent and give it to m
	m.Parent = n.Parent
}

func (t *RedBlackTree[T, V]) minimum(node *RedBlackTreeNode[T, V]) *RedBlackTreeNode[T, V] {
	for node.Left != t.NIL {
		node = node.Left
	}
	return node
}
