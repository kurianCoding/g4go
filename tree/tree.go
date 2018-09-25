package tree

type Node struct {
	Left   *Node // pointer to left node
	Right  *Node // pointer to right node
	Val    int   // value stored in this node
	Parent *Node
}

/*tree is given by its root node*/
type Tree struct {
	Name string
	Root *Node
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

// get root returns the root node
func (t *Tree) GetRoot() *Node {
	return t.Root
}
