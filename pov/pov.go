package pov

import "fmt"

// Define the Graph type here.
type Graph struct {
	value    string
	children []*node
}

type node struct {
	value    string
	children []*node
}

// Creates a new graph
func New() *Graph {
	g := new(Graph)
	g.children = make([]*node, 0)
	return g
}

// Returns a node with the value
func NewNode(nodeLabel string) *node {
	n := new(node)
	n.value = nodeLabel
	n.children = make([]*node, 0)
	return n
}

// Adds leaf nodes to the graph
func (g *Graph) AddNode(nodeLabel string) *node {
	new := NewNode(nodeLabel)
	g.children = append(g.children, new)
	return new // Returns memory address of new node
}

// Finds if a node exists in the graph
func (g *Graph) FindNode(nodeLabel string) *node {
	for i, n := range g.children {
		if g.children[i].value == nodeLabel {
			return n
		}
	}
	return nil // If not found
}

// Constructs tree from bottom up
// to: a node that has already been added
// from: a node to add
func (g *Graph) AddArc(from, to string) {
	// Add from if not exists in graph
	fromNode := g.FindNode(from)
	if fromNode == nil {
		new := g.AddNode(from)
		new.children = append(new.children, g.FindNode(to)) // Add to
	} else { // from already exists in graph
		fromNode.children = append(fromNode.children, g.FindNode(to)) // Add to
	}
}

// Dump method
// Returns a list of all arcs in the graph
// from -> to
// Test program will sort the list (we don't need to)
func (g *Graph) ArcList() []string {
	arcs := make([]string, 0)
	for _, n := range g.children {
		for _, c := range n.children {
			arcs = append(arcs, fmt.Sprintf("%s -> %s", n.value, c.value))
		}
	}
	return arcs
}

// Change the root of the graph
// Can create a new graph if we like or do the change in place
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {

	newGraph := New()

	fmt.Println("oldRoot:", oldRoot)
	fmt.Println("newRoot:", newRoot)

	return newGraph
}
