package tree

import "fmt"

/* add a node to the nearest empty leaf*/
func AddNode(n *Node, key *Node) error {
	var queue = make([]*Node, 0)
	for len(queue) > 0 {
		k := queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]
		if k.Left == nil {
			k.Left = key
			return nil
		} else {
			queue = append(queue, k.Left)
		}
		if k.Right == nil {
			k.Right = key
			return nil
		} else {
			queue = append(queue, k.Right)
		}

	}
	return fmt.Errorf("no empty leaf found")
}
