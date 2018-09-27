package tree

import "testing"

func TestAddNode(t *testing.T) {
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
