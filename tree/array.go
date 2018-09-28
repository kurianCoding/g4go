package tree

/* to add a value to the left of root*/
func AddLeft(key, node int) {
	if tree[node] == nil {
		panic("parent is empty")
	}
	tree[node*2] = key
	return
}

/* to add a value to the right of root*/
func AddRight(key, node int) {
	if tree[node] == nil {
		panic("parent is empty")
	}
	tree[node*2+1] = key
	return
}
