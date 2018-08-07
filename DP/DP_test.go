package DP

import "testing"

func TestHCF(t *testing.T) {
	hcf := HCF(4, 6)
	if hcf != 2 {
		t.Errorf("HCF testing failed got %d", hcf)
	}

}
func TestLenProf(t *testing.T) {
	maxCost := LenProf([]int{1, 5, 8, 9, 10, 17, 17, 20}, 8)
	if maxCost != 22 {
		t.Errorf("LenProf testing failed got %d", maxCost)
	}
}
func TestFib(t *testing.T) {
	fib := Fib(3)
	if fib != 2 {
		t.Errorf("Fibonacci testing failed got %d", fib)
	}
}

func TestCompleteSum(t *testing.T) {
	exists := CompleteSum([]int{3, 34, 4, 12, 5, 2}, 9)
	if !exists {
		t.Errorf("CompleteSum testing failed got %v", exists)
	}
}

func TestMinJumps(t *testing.T) {
	minimumJumps := MinJumps([]int{1, 3, 6, 3, 2, 3, 6, 8, 9, 5})
	if minimumJumps != 4 {
		t.Errorf("Minimum Jumps failed got %d", minimumJumps)
	}
}
func TestLimitRep(t *testing.T) {
	clearedRep := LimitRep([]int{1, 2, 3, 1, 2, 1, 2, 3}, 2)
	result := []int{1, 2, 3, 1, 2, 3}
	for key, val := range clearedRep {
		if val != result[key] {
			t.Errorf("Limit repetition falied")
		}
	}
}
func TestCleanZeros(t *testing.T) {
	cleanZeros := CleanZeros([]int{1, 0, 1, 2, 0, 1, 3, 0})
	clResult := []int{1, 1, 2, 1, 3, 0, 0, 0}
	for key, val := range cleanZeros {
		if val != clResult[key] {
			t.Errorf("Clean Zeros failed")
		}
	}
}

func TestRotateArray(t *testing.T) {
	rotatedArray := RotateArray([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rAresult := []int{5, 6, 7, 1, 2, 3, 4}
	for key, val := range rotatedArray {
		if val != rAresult[key] {
			t.Errorf("Rotate Array failed")
		}
	}
}

func TestCountChange(t *testing.T) {
	coin := []int{1, 2, 3}
	countc := countChange(coin, 3, 4)
	if countc != 4 {
		t.Errorf("count change failed %d needed got %d", 4, countc)
	}
}

func TestRecSum(t *testing.T) {
	set := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a, b := getSumArray(set, 15)
	if a != 0 || b != 5 {
		t.Errorf("getSumArray %d,%d expected, found %d,%d", 1, 3, a, b)
	}
}

//func TestPlusOne(t *testing.T) {
//	resultPlusOne := PlusOne([]int{1, 9, 9, 9, 9, 9, 9})
//	rPOresult := []int{2, 0, 0, 0, 0, 0, 0}
//	for key, val := range rPOresult {
//		if val != resultPlusOne[key] {
//			t.Errorf("Plus One failed test")
//		}
//	}
//}

/*flatteing of arrays into a single array, the problem is that you cant
  range over arrays, so it gets defined as a slice, when defined like that
  you can't pass values as slice of int*/

//func TestFlatten(t *testing.T) {
//	flatArray := Flatten([]int{[]int{1, 2}, []int{3, 4}})
//	resultArray := []int{1, 2, 3, 4}
//	for key, val := range resultArray {
//		if val != flatArray[key] {
//			t.Errorf("Flatten failed test")
//		}
//	}
//}
func TestKAnagrams(t *testing.T) {
	isAnagram := KAnagrams("grammar", "anagram", 3)
	if !isAnagram {
		t.Errorf(" K-Anagram failed test")
	}
}

func TestIsSquare(t *testing.T) {
	isSquare := IsSquare([]string{"ball", "area", "lead", "lady"})
	if !isSquare {
		t.Errorf(" Check Square failed test")
	}
}

/*func TestEditDistance(t *testing.T) {
	distance := EditDistance("sunday", "saturday")
	if distance != 3 {
		t.Errorf(" Edit distance has failed test expected 3 got %d", distance)
	}
}*/

func TestOptRate(t *testing.T) {
	priceArray := []int{1, 5, 8, 9, 10, 17, 17, 20}
	rodLength := len(priceArray)
	optimumPrice := OptRate(priceArray, rodLength)
	if optimumPrice != 22 {
		t.Errorf(" Optimum price has failed test expected 22 got %d", optimumPrice)
	}
}
func TestMapPall(t *testing.T) {
	freq := mapPall(112211)
	if freq != 1 {
		t.Errorf(" Optimum price has failed test expected %d got %d", 1, freq)
	}
}
