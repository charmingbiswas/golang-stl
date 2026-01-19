package trees

import (
	"cmp"
	"testing"
)

// Helper function to verify red-black tree properties
func verifyRedBlackProperties[T cmp.Ordered, V any](t *testing.T, tree *RedBlackTree[T, V]) {
	// Property 1: Root must be black
	if tree.Root != tree.NIL && tree.Root.NodeColor != BLACK {
		t.Error("Property violation: Root must be black")
	}

	// Verify other properties using helper
	verifyNode(t, tree, tree.Root)
}

func verifyNode[T cmp.Ordered, V any](t *testing.T, tree *RedBlackTree[T, V], node *RedBlackTreeNode[T, V]) int {
	if node == tree.NIL {
		return 1 // NIL nodes are black and count as 1 black node
	}

	// Property 3: Red nodes must have black children
	if node.NodeColor == RED {
		if node.Left.NodeColor == RED {
			t.Errorf("Property violation: Red node %v has red left child", node.Key)
		}
		if node.Right.NodeColor == RED {
			t.Errorf("Property violation: Red node %v has red right child", node.Key)
		}
	}

	// Property 4: All paths must have same number of black nodes
	leftBlackHeight := verifyNode(t, tree, node.Left)
	rightBlackHeight := verifyNode(t, tree, node.Right)

	if leftBlackHeight != rightBlackHeight {
		t.Errorf("Property violation: Black height mismatch at node %v (left: %d, right: %d)",
			node.Key, leftBlackHeight, rightBlackHeight)
	}

	// Add 1 if current node is black
	if node.NodeColor == BLACK {
		return leftBlackHeight + 1
	}
	return leftBlackHeight
}

func TestNewRedBlackTree(t *testing.T) {
	tree := NewRedBlackTree[int, string]()

	if tree == nil {
		t.Fatal("NewRedBlackTree returned nil")
	}

	if tree.Root != tree.NIL {
		t.Error("New tree root should be NIL")
	}

	if tree.NIL.NodeColor != BLACK {
		t.Error("NIL node should be black")
	}

	if tree.Size() != 0 {
		t.Errorf("New tree size should be 0, got %d", tree.Size())
	}

	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}
}

func TestInsertSingle(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	tree.Insert(10, "ten")

	if tree.Root.Key != 10 {
		t.Errorf("Expected root key 10, got %d", tree.Root.Key)
	}

	if tree.Root.Value != "ten" {
		t.Errorf("Expected root value 'ten', got %s", tree.Root.Value)
	}

	if tree.Root.NodeColor != BLACK {
		t.Error("Root should be black")
	}

	if tree.Size() != 1 {
		t.Errorf("Expected size 1, got %d", tree.Size())
	}

	verifyRedBlackProperties(t, tree)
}

func TestInsertMultiple(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	values := []struct {
		key   int
		value string
	}{
		{10, "ten"},
		{5, "five"},
		{15, "fifteen"},
		{3, "three"},
		{7, "seven"},
		{12, "twelve"},
		{17, "seventeen"},
	}

	for _, v := range values {
		tree.Insert(v.key, v.value)
		verifyRedBlackProperties(t, tree)
	}

	if tree.Size() != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), tree.Size())
	}

	// Verify all values can be found
	for _, v := range values {
		node, found := tree.Search(v.key)
		if !found {
			t.Errorf("Key %d not found", v.key)
		}
		if node.Value != v.value {
			t.Errorf("Expected value %s for key %d, got %s", v.value, v.key, node.Value)
		}
	}
}

func TestInsertDuplicate(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	tree.Insert(10, "ten")
	tree.Insert(10, "updated_ten")

	if tree.Size() != 1 {
		t.Errorf("Expected size 1 after duplicate insert, got %d", tree.Size())
	}

	node, found := tree.Search(10)
	if !found {
		t.Fatal("Key 10 not found")
	}

	if node.Value != "updated_ten" {
		t.Errorf("Expected updated value 'updated_ten', got %s", node.Value)
	}

	verifyRedBlackProperties(t, tree)
}

func TestInsertSequential(t *testing.T) {
	tree := NewRedBlackTree[int, int]()

	// Insert sequential values (worst case for unbalanced BST)
	for i := 1; i <= 10; i++ {
		tree.Insert(i, i*10)
		verifyRedBlackProperties(t, tree)
	}

	if tree.Size() != 10 {
		t.Errorf("Expected size 10, got %d", tree.Size())
	}

	// Verify all values
	for i := 1; i <= 10; i++ {
		node, found := tree.Search(i)
		if !found {
			t.Errorf("Key %d not found", i)
		}
		if node.Value != i*10 {
			t.Errorf("Expected value %d for key %d, got %d", i*10, i, node.Value)
		}
	}
}

func TestSearch(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	tree.Insert(10, "ten")
	tree.Insert(5, "five")
	tree.Insert(15, "fifteen")

	// Search existing keys
	node, found := tree.Search(10)
	if !found || node.Value != "ten" {
		t.Error("Failed to find key 10")
	}

	node, found = tree.Search(5)
	if !found || node.Value != "five" {
		t.Error("Failed to find key 5")
	}

	// Search non-existing key
	_, found = tree.Search(100)
	if found {
		t.Error("Should not find non-existing key 100")
	}
}

func TestSearchEmptyTree(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	_, found := tree.Search(10)
	if found {
		t.Error("Should not find key in empty tree")
	}
}

func TestDeleteSingle(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	tree.Insert(10, "ten")

	deleted := tree.Delete(10)
	if !deleted {
		t.Error("Failed to delete key 10")
	}

	if !tree.IsEmpty() {
		t.Error("Tree should be empty after deleting only element")
	}

	if tree.Size() != 0 {
		t.Errorf("Expected size 0, got %d", tree.Size())
	}
}

func TestDeleteNonExisting(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	tree.Insert(10, "ten")

	deleted := tree.Delete(100)
	if deleted {
		t.Error("Should not delete non-existing key")
	}

	if tree.Size() != 1 {
		t.Errorf("Size should remain 1, got %d", tree.Size())
	}
}

func TestDeleteMultiple(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	keys := []int{10, 5, 15, 3, 7, 12, 17, 1, 4, 6, 8}

	// Insert all keys
	for _, key := range keys {
		tree.Insert(key, "value")
	}

	// Delete half the keys
	keysToDelete := []int{3, 7, 15, 1, 8}
	for _, key := range keysToDelete {
		deleted := tree.Delete(key)
		if !deleted {
			t.Errorf("Failed to delete key %d", key)
		}
		verifyRedBlackProperties(t, tree)

		// Verify key is gone
		_, found := tree.Search(key)
		if found {
			t.Errorf("Key %d should not be found after deletion", key)
		}
	}

	expectedSize := len(keys) - len(keysToDelete)
	if tree.Size() != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, tree.Size())
	}

	// Verify remaining keys still exist
	remainingKeys := []int{10, 5, 12, 17, 4, 6}
	for _, key := range remainingKeys {
		_, found := tree.Search(key)
		if !found {
			t.Errorf("Key %d should still exist", key)
		}
	}
}

func TestDeleteAll(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	keys := []int{10, 5, 15, 3, 7, 12, 17}

	// Insert all keys
	for _, key := range keys {
		tree.Insert(key, "value")
	}

	// Delete all keys
	for _, key := range keys {
		deleted := tree.Delete(key)
		if !deleted {
			t.Errorf("Failed to delete key %d", key)
		}
		if tree.Root != tree.NIL {
			verifyRedBlackProperties(t, tree)
		}
	}

	if !tree.IsEmpty() {
		t.Error("Tree should be empty after deleting all elements")
	}

	if tree.Size() != 0 {
		t.Errorf("Expected size 0, got %d", tree.Size())
	}
}

func TestIterator(t *testing.T) {
	tree := NewRedBlackTree[int, string]()
	expected := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}

	for k, v := range expected {
		tree.Insert(k, v)
	}

	// Verify iterator returns values in sorted order
	lastKey := 0
	count := 0
	for key, value := range tree.ForwardIterator() {
		if key <= lastKey && count > 0 {
			t.Error("Iterator should return keys in sorted order")
		}
		if expected[key] != value {
			t.Errorf("Expected value %s for key %d, got %s", expected[key], key, value)
		}
		lastKey = key
		count++
	}

	if count != len(expected) {
		t.Errorf("Expected %d iterations, got %d", len(expected), count)
	}
}

func TestIteratorEmpty(t *testing.T) {
	tree := NewRedBlackTree[int, string]()

	count := 0
	for range tree.ForwardIterator() {
		count++
	}

	if count != 0 {
		t.Errorf("Empty tree iterator should yield no values, got %d", count)
	}
}

func TestStringKeys(t *testing.T) {
	tree := NewRedBlackTree[string, int]()
	tree.Insert("apple", 1)
	tree.Insert("banana", 2)
	tree.Insert("cherry", 3)

	node, found := tree.Search("banana")
	if !found || node.Value != 2 {
		t.Error("Failed to find string key 'banana'")
	}

	deleted := tree.Delete("apple")
	if !deleted {
		t.Error("Failed to delete string key 'apple'")
	}

	verifyRedBlackProperties(t, tree)
}

func TestLargeDataSet(t *testing.T) {
	tree := NewRedBlackTree[int, int]()

	// Insert 1000 elements
	for i := 0; i < 1000; i++ {
		tree.Insert(i, i*2)
	}

	if tree.Size() != 1000 {
		t.Errorf("Expected size 1000, got %d", tree.Size())
	}

	verifyRedBlackProperties(t, tree)

	// Delete every other element
	for i := 0; i < 1000; i += 2 {
		tree.Delete(i)
	}

	if tree.Size() != 500 {
		t.Errorf("Expected size 500 after deletions, got %d", tree.Size())
	}

	verifyRedBlackProperties(t, tree)

	// Verify remaining elements
	for i := 1; i < 1000; i += 2 {
		node, found := tree.Search(i)
		if !found || node.Value != i*2 {
			t.Errorf("Failed to find key %d or incorrect value", i)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tree := NewRedBlackTree[int, string]()

	if !tree.IsEmpty() {
		t.Error("New tree should be empty")
	}

	tree.Insert(1, "one")
	if tree.IsEmpty() {
		t.Error("Tree with elements should not be empty")
	}

	tree.Delete(1)
	if !tree.IsEmpty() {
		t.Error("Tree should be empty after deleting all elements")
	}
}

func TestComplexInsertDeleteSequence(t *testing.T) {
	tree := NewRedBlackTree[int, string]()

	// Complex sequence of operations
	operations := []struct {
		op    string
		key   int
		value string
	}{
		{"insert", 50, "fifty"},
		{"insert", 25, "twenty-five"},
		{"insert", 75, "seventy-five"},
		{"insert", 10, "ten"},
		{"insert", 30, "thirty"},
		{"delete", 25, ""},
		{"insert", 60, "sixty"},
		{"insert", 80, "eighty"},
		{"delete", 50, ""},
		{"insert", 5, "five"},
		{"delete", 10, ""},
		{"insert", 70, "seventy"},
	}

	expectedSize := 0
	for _, op := range operations {
		if op.op == "insert" {
			tree.Insert(op.key, op.value)
			expectedSize++
		} else {
			tree.Delete(op.key)
			expectedSize--
		}
		verifyRedBlackProperties(t, tree)
	}

	if tree.Size() != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, tree.Size())
	}
}
