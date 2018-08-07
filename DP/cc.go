package DP

/*
   we have a set of currencies, from which to make change
   this function ouputs the number of ways this can be
   done
*/

func countChange(change []int, length, val int) int {
	if val == 0 {
		// value reached
		return 1
	}
	if val < 0 {
		// value exhausted and overshot
		return 0
	}
	if length <= 0 && val >= 1 {
		// change was not reached, even after exhausting all the coinage
		return 0
	}
	return countChange(change, length-1, val) + countChange(change, length, val-change[length-1])
}
