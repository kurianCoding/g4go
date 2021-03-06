package tree

import "testing"

func TestInOrderTraversal(t *testing.T) {

	//root := NewNode(10)
	//key := NewNode(15)
	//key2 := NewNode(20)
	//key3 := NewNode(30)
	//root.AddLeft(key)
	//root.AddRight(key2)
	//key2.AddRight(key3)
	//[> 	     root
	//key       key2
	//keye3
	//*/
	//InOrderTraversal(root)
}

func TestBalanced(t *testing.T) {
	root := NewNode(10)
	key := NewNode(15)
	key2 := NewNode(20)
	key3 := NewNode(30)
	root.AddLeft(key)
	root.AddRight(key2)
	key2.AddRight(key3)
	/* 	     root
		key       key2
	    keye3
	*/
	if num := Balanced(root); num != true {
		t.Errorf("expected %v got %v", true, num)
	}
}

func TestSize(t *testing.T) {
	root := NewNode(10)
	key := NewNode(15)
	key2 := NewNode(20)
	root.AddLeft(key)
	root.AddRight(key2)
	if num := Size(root); num != 3 {
		t.Errorf("expected %d got %d", 3, num)
	}
}

func TestDeleteNode(t *testing.T) {
	root := NewNode(10)
	key := NewNode(15)
	key2 := NewNode(20)
	root.AddLeft(key)
	root.AddRight(key2)
	DeleteNode(root, key)
}

func TestAddNode(t *testing.T) {
	root := NewNode(10)
	key := NewNode(15)
	a := NewNode(20)
	root.AddLeft(a)
	err := AddNode(root, key)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestCatlan(t *testing.T) {
	n := Catlan(4)
	if n != 14 {
		t.Errorf("expected %d got %d", 14, n)
	}
	n = PCatlan(4)
	if n != 14 {
		t.Errorf("expected %d got %d", 14, n)
	}
}

func TestSumAndProduct(t *testing.T) {
	n1 := NewNode(10)
	n2 := NewNode(4)
	n3 := NewNode(9)
	n4 := NewNode(11)
	n5 := NewNode(1)
	n6 := NewNode(8)
	n7 := NewNode(6)
	n8 := NewNode(7)
	n9 := NewNode(2)
	n9.AddRight(n7)
	n9.AddLeft(n8)
	n8.AddRight(n6)
	n8.AddLeft(n5)
	n7.AddRight(n4)
	n7.AddLeft(n3)
	n5.AddLeft(n2)
	n5.AddRight(n1)
	sum := SumAndProduct(n9)
	if sum != 208 {
		t.Errorf("expected %d got %d", 208, sum)
	}

}
