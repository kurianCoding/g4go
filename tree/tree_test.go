package tree

import "testing"

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
