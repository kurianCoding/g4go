package divncc

import "testing"

func TestNthCmb(t *testing.T) {
	res := NthCmb([]int{2, 3, 6, 7, 9}, []int{1, 4, 8, 10}, 5)
	if res != 6 {
		t.Errorf("expected %d, got %d", 6, res)
	}
}

func TestIsMajority(t *testing.T) {
	res := IsMajority([]int{1, 2, 3, 3, 3, 3, 10}, 3)
	if !res {
		t.Errorf("expected %v, got %v", true, res)
	}
}

func TestCountRot(t *testing.T) {
	if v := CountRot([]int{12, 15, 18, 2, 3, 6}); v != 3 {
		t.Errorf("expected %d, got %d", 2, v)
	}
}

func TestExtraElement(t *testing.T) {
	if v := ExtraElement([]int{1, 2, 4}, []int{1, 2, 3, 4}); v != 3 {
		t.Errorf("expected %d, got %d", 3, v)
	}
}
