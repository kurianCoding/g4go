package backtrack

import (
	"testing"
)

func TestBrackets(t *testing.T) {
	cleanStrings := EleminateBrackets("()())()")
	if len(cleanStrings) != 3 {
		t.Errorf("expected %d got %d", 3, len(cleanStrings))
	}
}

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
func TestHamiltonianCycle(t *testing.T) {
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
func TestSudoku(t *testing.T) {
	a := [][]int{{3, 0, 6, 5, 0, 8, 4, 0, 0},
		{5, 2, 0, 0, 0, 0, 0, 0, 0},
		{0, 8, 7, 0, 0, 0, 0, 3, 1},
		{0, 0, 3, 0, 1, 0, 0, 8, 0},
		{9, 0, 0, 8, 6, 3, 0, 0, 5},
		{0, 5, 0, 0, 9, 0, 6, 0, 0},
		{1, 3, 0, 0, 0, 0, 2, 5, 0},
		{0, 0, 0, 0, 0, 0, 0, 7, 4},
		{0, 0, 5, 2, 0, 6, 3, 0, 0}}

	solution := SD(a)
	if !solution {
		t.Errorf("expected %v got %v", true, false)
	}
}
