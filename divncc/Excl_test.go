package divncc

import (
	"github.com/fortytw2/leaktest"
	"testing"
)

func TestIncOddNum(t *testing.T) {
	defer leaktest.Check(t)()
	res := OddNum([]int{1, 1, 2, 4, 4, 5, 5, 6, 6})
	if res != 2 {
		t.Errorf("expected %d, got %d", 2, res)
	}
}
func TestIncNthCmb(t *testing.T) {
	res := NthCmb([]int{2, 3, 6, 7, 9}, []int{1, 4, 8, 10}, 5)
	if res != 6 {
		t.Errorf("expected %d, got %d", 6, res)
	}
}
