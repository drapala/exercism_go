package main

import (
	"fmt"
	"reflect"
)

type Record struct {
	ID     int
	Parent int // 0 if parent
}

type Node struct {
	ID       int
	Children []*Node
}

func sortRecord(records []Record) []Record {
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
	return records
}

func appendToChild(OutputNode *Node, record Record) {
	fmt.Println("Appending to child")
	fmt.Println(OutputNode, record)
	OutputNode.Children = append(OutputNode.Children, &Node{ID: record.ID})	
}

func findParentNode(PreviousNodes *Node, Parent int) (*Node, error){
	fmt.Println("Finding parent node ", Parent)
	fmt.Println("Previous Nodes: ", PreviousNodes)
	for _, node := range PreviousNodes.Children {
		fmt.Println("Checking Node: ", node)
		// Parent is at this level
		if node.ID == Parent {
			fmt.Println("Found parent node: ", node.ID)
			return node, nil
		} else {
			// Parent is at a lower level
			ParentNode, err := findParentNode(node, Parent)
			if err == nil {
				return ParentNode, nil
			}
		}
	}
	return nil, fmt.Errorf("Parent node not found!")
}

func printNodesRecursive(node *Node) {
	fmt.Println(node.ID)
	for _, child := range node.Children {
		printNodesRecursive(child)
	}
}

func Build(records []Record) (*Node, error) {
	// Edge case handling
	if reflect.DeepEqual(records, []Record{}) {
		return nil, nil
	}
	
	// Sort the records
	// sortRecord(records)
	fmt.Println(sortRecord(records))

	OutputNode := Node{}
	// Loop over records
	for _, record := range records {
		// Root node
		if record.Parent == 0 && record.ID == 0 {
			OutputNode.ID = record.ID
		} else {
			// Children belong at current level
			if record.Parent == OutputNode.ID {
				appendToChild(&OutputNode, record)
			} else {
				// Children must belong at a lower level
				ParentNode, error := findParentNode(&OutputNode, record.Parent)
				if error == nil {
					fmt.Println("Trying to append: ", record)
					appendToChild(ParentNode, record)
				} else {
					panic("Parent node not found!")
				}
			}
		}
	}
	// Remover after
	printNodesRecursive(&OutputNode)

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
		{ID: 2, Parent: 1},
		{ID: 1, Parent: 0},
		{ID: 3, Parent: 2},
		{ID: 0},
	}

	OutputNode, err := Build(InputRecord)

	if err != nil {
		fmt.Println(err)
	} else {
		printNodesRecursive(OutputNode)
	}
}