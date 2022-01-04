package binarysearchtree

type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// NewBst creates and returns a new SearchTreeData.
func NewBst(i int) SearchTreeData {
	return SearchTreeData{left: nil, data: i, right: nil}
}

// Insert inserts an int into the SearchTreeData.
// Inserts happen based on the rules of a BinarySearchTree
func (std *SearchTreeData) Insert(i int) {
	newNode := NewBst(i) // Create new node
	var inserted bool    // Flag for successful insertion
	node := std          // Traverse from head

	for !inserted {
		// Check if we go left or right
		if i <= node.data { // Go left
			if node.left == nil { // Found an empty node to insert to
				node.left = &newNode // insert
				inserted = true
			} else {
				node = node.left // Traverse left
			}
		} else { // Go right
			if node.right == nil {
				node.right = &newNode
				inserted = true
			} else {
				node = node.right
			}
		}
	}
}

// ========================================================
// Traversal: See book CTCI - Page 103 (book page, not pdf)
// void inOrderTraversal(TreeNode node) {
// 	if (node != null) {
// 		inOrderTraversal(node.left);
// 		visit(node);
// 		inOrderTraversal(node.right);
// 	}
// }
// ========================================================

func inOrderStringAppend(Node *SearchTreeData, result *[]string, f func(int) string) {
	if Node != nil {
		inOrderStringAppend(Node.left, result, f)  // Traverse left
		*result = append(*result, f(Node.data))    // Append to slice
		inOrderStringAppend(Node.right, result, f) // Traverse right
	}
}

func inOrderIntAppend(Node *SearchTreeData, result *[]int, f func(int) int) {
	if Node != nil {
		inOrderIntAppend(Node.left, result, f)  // Traverse left
		*result = append(*result, f(Node.data)) // Append to slice
		inOrderIntAppend(Node.right, result, f) // Traverse right
	}
}

// MapString returns the ordered contents of SearchTreeData as a []string.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []string ["1", "3", "5", "7"].
func (std *SearchTreeData) MapString(f func(int) string) (result []string) {
	result = make([]string, 0)
	inOrderStringAppend(std, &result, f) // Pointer to slice
	return result
}

// MapInt returns the ordered contents of SearchTreeData as an []int.
// The values are in increasing order starting with the lowest int value.
// SearchTreeData that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (std *SearchTreeData) MapInt(f func(int) int) (result []int) {
	result = make([]int, 0)
	inOrderIntAppend(std, &result, f) // Pointer to slice
	return result
}
