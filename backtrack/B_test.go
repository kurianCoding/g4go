package backtrack

import (
	"testing"
)

func TestKnightsTour(t *testing.T) {
	exist := KnightsTour(0, 0)
	if !exist {
		t.Errorf(" expected %v got %v", true, false)
	}
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
