package linkedlist

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	Val interface{}
	next *Node
	prev *Node
}

var ErrEmptyList = fmt.Errorf("empty list")

// ... is for Variadic functions
// https://go.dev/ref/spec#Passing_arguments_to_..._parameters
func NewList(args ...interface{}) *List {
	l := List{} // Create new list
	var prevNode *Node
	// Loop over items
	for i, v := range args {
		currNode := Node{Val: v} // Create node with value
		if i == 0 { // If first node, set prev to nil, and update list head
			// Current Node ops:
			currNode.prev = nil
			// currNode.next in next loop prevNode.next

			// Previous Node ops: 
			// None

			// List ops
			l.head = &currNode // set head
		} else if i == len(args) - 1 { // If last node, set next to nil, and update list tail
			// Current Node ops:
			currNode.prev = prevNode
			currNode.next = nil

			// Previous Node ops: 
			prevNode.next = &currNode

			// List ops
			l.tail = &currNode // set tail
		} else { // In the middle somewhere			
			// Current Node ops:
			currNode.prev = prevNode
			// currNode.next in next loop prevNode.next

			// Previous Node ops: 
			prevNode.next = &currNode

			// List ops
			// None	
		}
		
		// Edge case handling
		if len(args) == 1 { // Single entry
			// Current Node ops:
			// No currNode.prev
			currNode.next = nil

			// Previous Node ops: 
			// None

			// List ops
			l.tail = &currNode // set tail			
		}
		prevNode = &currNode // Store memory address of Curent Node for next loop run
	}
	return &l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	fmt.Println("value:", v)
	// Create a newNode with value
	newNode := Node{Val: v}
	fmt.Println("newNode 1:", newNode)
	// Set newNode.next = l.head
	newNode.next = l.head
	fmt.Println("head 1:", l.head)
	fmt.Println("newNode 2:", newNode)
	// Set newNode.prev = nil
	newNode.prev = nil
	// Set l.head.prev = &newNode
	if l.head == nil { // empty list, append in
		l.head = &newNode
	} else { // Update previous to point to newNode
		l.head.prev = &newNode
		l.head = &newNode
	}
	fmt.Println("head 2:", l.head)
	if l.head.next != nil {
		fmt.Println("NodeAfter val:", l.head.next.Val)
	}
	fmt.Println("next!")
}

func (l *List) PushBack(v interface{}) {
	panic("Please implement the PushBack function")
}

func (l *List) PopFront() (interface{}, error) {
	panic("Please implement the PopFront function")
}

func (l *List) PopBack() (interface{}, error) {
	panic("Please implement the PopBack function")
}

func (l *List) Reverse() {
	// Switch head and tail
	tempNode := l.tail
	l.tail = l.head
	l.head = tempNode

	// Switch head and tail per node
	currNode := l.head
	for currNode != nil {
		temp := currNode.next
		currNode.next = currNode.prev
		currNode.prev = temp
		currNode = currNode.next
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
