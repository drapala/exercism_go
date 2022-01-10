package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden struct {
	lookup map[string][]string
}

type Child struct {
	name   string
	plants []string
}

var plantMap = map[string]string{
	"V": "violets",
	"R": "radishes",
	"C": "clover",
	"G": "grass",
}

// Initiates a new Garden and returns the pointer.
// Error if any inputs are incorrect.
//
// Notes:
// 1. children could be out of order, but diagram is always in order

func NewGarden(diagram string, children []string) (*Garden, error) {
	// Sort children without modifying the original slice
	children_sorted := make([]string, len(children)) // make a new slice with same length
	_ = copy(children_sorted, children)              // Copy to new slice
	sort.Strings(children_sorted)                    // Sort in place
	// Decipher diagram
	rows := strings.Split(diagram, "\n") // Break up the diagram into rows
	rows = rows[1:]                      // Remove the first row which is a newline
	// ########################################
	// Error handling
	// ########################################
	// Deal with wrong diagram format
	// If first character in diagram is not newline, then error
	if diagram[0] != '\n' {
		return nil, fmt.Errorf("diagram must start with a newline")
	}
	// ########################################
	// Deal with mismatched rows
	// Deal with odd no. of cups
	// ########################################
	if len(rows)%2 != 0 || len(rows[0])%2 != 0 || len(rows[1])%2 != 0 || len(rows[0]) != len(rows[1]) {
		return nil, fmt.Errorf("rows are badly formatted")
	}
	// ########################################
	// Loop over children_sorted and create plants - append to garden
	var g Garden
	g.lookup = make(map[string][]string) // Initialize map
	for i, c := range children_sorted {
		var child Child // Initialize every loop
		child.name = c
		// Append the plants for this child to the plants array
		for j := range rows {
			// ########################################
			// Deal with invalid cup codes
			// ########################################
			plant1, ok1 := plantMap[rows[j][2*i:2*i+1]]
			plant2, ok2 := plantMap[rows[j][2*i+1:2*i+2]]
			if !ok1 || !ok2 { // If plant is not in map
				return nil, fmt.Errorf("invalid cup code")
			} else {
				child.plants = append(child.plants, plant1)
				child.plants = append(child.plants, plant2)
			}
		}
		// ########################################
		// Deal with duplicate child names
		// ########################################
		_, ok := g.lookup[child.name]
		if ok { // Insert into map only if child not present
			return nil, fmt.Errorf("duplicated child name")
		} else {
			g.lookup[child.name] = child.plants // Add the child to map based on child name as unique key
		}
	}
	return &g, nil
}

// Method takes in name of child and returns the plants for that child as an array
// The boolean indicates whether the child was found in the garden
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.lookup[child]
	return plants, ok
}
