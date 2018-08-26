package backtrack

import (
	"testing"
)

func TestKnightsTour(t *testing.T) {
	//exist := KnightsTour(0, 0)
	//if !exist {
	//t.Errorf(" expected %v got %v", true, false)
	//}
}

func TestRatInMaze(t *testing.T) {
	exist := RatInMaze([][]int{
		{1, 0, 0, 0},
		{1, 1, 0, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 1},
	}, 0, 0)
	if !exist {
		t.Errorf(" expected %v got %v", true, false)
	}
}
func TestHM(t *testing.T) {
	A := [][]bool{
		{false, true, false, true, false},
		{true, false, true, true, true},
		{false, true, false, false, true},
		{true, true, false, false, false},
		{false, true, true, false, false},
	}
	exist := hasHM(A, 5)
	if exist {
		t.Errorf("expected %v got %v", false, true)
	}
}
