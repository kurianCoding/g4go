package tree

/*
   transverse a tree without having to allocate
   extra space or recursion. The idea is to push
   the current node as the right subtree of the
   rightmost substree of the current left subtree

   so, you check if the tree has left and right nodes
   if it has a righ node, search for the right most
   subtree of the left subtree, i.e. the point where
   right subtree is nil.(which means its a leaf). At this
   point assign the right subtree to the found node.
   continue by switching to the left subtree. At the end
   you will have re-made the tree
*/
func MorrisTraversal(root *Node) {}
