package tree

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

// Checks if child exists already in the node's children
func childExists(node *Node, ID int) bool {
	for _, child := range node.Children {
		if child.ID == ID {
			return true
		}
	}
	return false
}

// Appends a child node to the parent node
func appendToParent(OutputNode *Node, record Record) (*Node, error) {
	// Check if child exists
	if childExists(OutputNode, record.ID) {
		return nil, fmt.Errorf("child already exists")
	} else {
		// If parent is higher than child, throw error
		if OutputNode.ID > record.ID {
			return nil, fmt.Errorf("parent is higher than child")
		}
		// If no errors, append to children
		ChildAddress := &Node{ID: record.ID}
		OutputNode.Children = append(OutputNode.Children, ChildAddress)
		return ChildAddress, nil
	}
}

// Returns the memory address of the Parent node from the cache map
func findParentNode(AddressMap map[int]*Node, Parent int) (*Node, error) {
	ParentNode := AddressMap[Parent]
	if ParentNode == nil {
		return nil, fmt.Errorf("parent node not found")
	} else {
		return ParentNode, nil
	}
}

// Builds a tree from a list of records
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

	// Create a slice to contain memory addresses of nodes
	var AddressMap map[int]*Node
	AddressMap = make(map[int]*Node)

	// Loop over records
	for i, record := range records {
		// This is the root node
		if record.Parent == 0 && record.ID == 0 {
			// Not first record, but is root
			if i != 0 && record.ID == 0 {
				return nil, fmt.Errorf("duplicate root")
			}
			// Create the root node
			OutputNode.ID = record.ID
			// 1B. Add the root node to the address map
			AddressMap[record.ID] = &OutputNode
		} else {

			// #######################################
			// Error handling - Parent = Root
			// #######################################
			// Root node has parent
			if record.Parent != 0 && record.ID == 0 {
				return nil, fmt.Errorf("root node has a parent")
			}
			// First record, but not root
			if i == 0 && record.ID != 0 {
				return nil, fmt.Errorf("first record is not root")
			}
			// #######################################
			// Children that belong at root level
			// #######################################
			if record.Parent == OutputNode.ID {
				ChildAddress, err := appendToParent(&OutputNode, record)
				AddressMap[record.ID] = ChildAddress
				if err != nil {
					return nil, fmt.Errorf("child already exists")
				}
			} else {
				// #######################################
				// Children lower than root
				// #######################################
				ParentNode, _ := findParentNode(AddressMap, record.Parent)
				if ParentNode == nil {
					return nil, fmt.Errorf("parent node doesn't exist")
				}
				// Parent found, append child to parent
				ChildAddress, err := appendToParent(ParentNode, record)
				if err != nil {
					return nil, fmt.Errorf("child already exists")
				}
				// Add valid child address to cache map for next iteration
				AddressMap[record.ID] = ChildAddress
			}
		}
	}

	return &OutputNode, nil
}
