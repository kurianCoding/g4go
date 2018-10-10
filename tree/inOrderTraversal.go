package tree

import "fmt"

/*
   a tree can be traversed from left to right we write a program
   to do this
*/
func InOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	var curr *Node
	curr = root
	q := make([]*Node, 0)
	q = append(q, curr)
	q = q[:len(q)-1]
	for len(q) > 0 || curr != nil { /*
			check if the queue holding nodes is none empty or
			if the current node is not pointing to a nil struct
		*/
		if curr != nil {
			for curr.Left != nil { /*
					check if the currend node is pointing to
					a nil vector
				*/
				curr = curr.Left
				q = append(q, curr) // push the left most nodes on top of queue
			}
		}
		curr = q[len(q)-1]
		if len(q)-1 > 0 {
			q = q[:len(q)-1]
		}

		fmt.Printf("%v ", curr.Val) // print the value of currnt node
		curr = curr.Right           // visit the right hand node after all the left nodes
	}
	return
}
