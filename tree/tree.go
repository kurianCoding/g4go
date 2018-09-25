package tree

type Node struct {
	Left  *Node // pointer to left node
	Right *Node // pointer to right node
	Val   int   // value stored in this node
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}
func (n *Node) AddLeft(m *Node) {
	n.Left = m
}
func (n *Node) AddRight(m *Node) {
	n.Right = m
}
