package main 

// struct for the node 
type Node struct {
	val string 
	next *Node
}

// constructor function to initialize a new node 
func NewNode(val string) *Node{
	return &Node{val: val, next: nil}
}

// method to set the next node 
func (n *Node) SetNext(node *Node){
	n.next = node 
}

// Stringer interface to support string representation
func (n *Node) String() string {
    return n.val
}