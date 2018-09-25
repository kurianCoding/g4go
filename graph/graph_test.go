package graph

import "testing"

func TestCountTrees(t *testing.T) {
	v := CountTrees([][]int{{0, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 1}})
	if v != 2 {
		t.Errorf("expected %d, got %d", 2, v)
	}
}
