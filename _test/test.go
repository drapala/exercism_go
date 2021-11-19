package main

import (
	"fmt"
	"reflect"
	"sort"
)

type Record struct {
	ID     int
	Parent int // 0 if root
}

type Node struct {
	ID       int
	Children []*Node
}

// Only runs once during execution
func sortRecord(records []Record) {
	// Sort records by Parent first, and then ID
	for i := 0; i < len(records); i++ {
		for j := i + 1; j < len(records); j++ {
			// If Current Node's Parent is greater than the next node's Parent
			// Then swap them
			if records[i].Parent > records[j].Parent {
				records[i], records[j] = records[j], records[i]
			}
			// If Parents are the same, but current node's ID is greater than the next node's ID
			// Then swap them
			if records[i].ID > records[j].ID && records[i].Parent == records[j].Parent {
				records[i], records[j] = records[j], records[i]
			}
		}
	}
}

// Only runs once during execution
func checkContinuousIDs(records []Record) bool {
	var ID_array []int
	for i := 0; i < len(records); i++ {
		ID_array = append(ID_array, records[i].ID)
	}
	// Sort array
	sort.Ints(ID_array)

	// Check if IDs are continuous
	for i := 0; i <= len(ID_array)-1; i++ {
		if i != 0 {
			if ID_array[i] > ID_array[i-1]+1 {
				return false
			}
		}
	}
	return true
}

func childExists(node *Node, ID int) bool {
	for _, child := range node.Children { // Room for optimization
		if child.ID == ID {
			return true
		}
	}
	return false
}

func appendToParent(OutputNode *Node, record Record) error {
	// Check if child exists
	if childExists(OutputNode, record.ID) {
		return fmt.Errorf("child already exists")
	} else {
		// If parent is higher than child, throw error
		if OutputNode.ID > record.ID {
			return fmt.Errorf("parent is higher than child")
		}
		// If no errors, append to children
		OutputNode.Children = append(OutputNode.Children, &Node{ID: record.ID})
		return nil
	}
}

func findParentNode(PreviousNodes *Node, Parent int) (*Node, error) {
	for _, node := range PreviousNodes.Children {
		// Parent is at this level
		if node.ID == Parent {
			return node, nil
		} else {
			// Parent is at a lower level
			ParentNode, err := findParentNode(node, Parent) // Room for optimization
			if err == nil {
				return ParentNode, nil
			}
		}
	}
	return nil, fmt.Errorf("parent node not found")
}

func Build(records []Record) (*Node, error) {
	// Edge case handling
	if reflect.DeepEqual(records, []Record{}) {
		return nil, nil
	}

	// Sort the records
	sortRecord(records)

	// Check if IDs are continuous
	if !checkContinuousIDs(records) {
		return nil, fmt.Errorf("IDs are not continuous")
	}

	// Create the root node
	OutputNode := Node{}

	// Loop over records
	for i, record := range records {
		// This is the root node
		if record.Parent == 0 && record.ID == 0 {
			// Not first record, but is root
			if i != 0 && record.ID == 0 {
				return nil, fmt.Errorf("duplicate root")
			}
			OutputNode.ID = record.ID
		} else {
			// Error handling for the root node cases
			// Root node has parent
			if record.Parent != 0 && record.ID == 0 {
				return nil, fmt.Errorf("root node has a parent")
			}
			// First record, but not root
			if i == 0 && record.ID != 0 {
				return nil, fmt.Errorf("first record is not root")
			}
			// Children belong at root level
			if record.Parent == OutputNode.ID {
				if appendToParent(&OutputNode, record) != nil {
					return nil, fmt.Errorf("child already exists")
				}
			} else {
				// Children must belong at a lower level of root
				ParentNode, error := findParentNode(&OutputNode, record.Parent)

				if error == nil {
					if appendToParent(ParentNode, record) != nil {
						return nil, fmt.Errorf("child already exists")
					}
				} else {
					return nil, fmt.Errorf("parent node doesn't exist")
				}
			}
		}
	}

	return &OutputNode, nil
}

func main() {
	// InputRecord := []Record{
	// 	{ID: 5, Parent: 1},
	// 	{ID: 3, Parent: 2},
	// 	{ID: 2, Parent: 0},
	// 	{ID: 4, Parent: 1},
	// 	{ID: 1, Parent: 0},
	// 	{ID: 0},
	// 	{ID: 6, Parent: 2},
	// }
	InputRecord := []Record{
		{ID: 5, Parent: 2},
		{ID: 3, Parent: 2},
		{ID: 2, Parent: 0},
		{ID: 4, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 0},
		{ID: 6, Parent: 2},
	}

	OutputNode, err := Build(InputRecord)

	if err != nil {
		fmt.Println(err)
	} else {
		printNodesRecursive(OutputNode)
	}
}
func printNodesRecursive(node *Node) {
	//fmt.Println(node.ID)
	for _, child := range node.Children {
		printNodesRecursive(child)
	}
}
