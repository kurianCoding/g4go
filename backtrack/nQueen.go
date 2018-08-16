package backtrack

/*
   Given a chess board, the idea is
   to see if one can place n Queens on
   a chess board without putting them
   in the line of attack of one antoher
   the queen is a pice that can attack
   all the squares in the same row and
   column and also attack all the squares
   which are diagonally in line with that
   square.


   the given function, takes in a value of n
   and returns a boolean true or false to
   say if such a formation is possible or not
*/
func qSafe(sol [][]int, a int, b int) bool {
	for i := 0; i < b; i++ {
		if sol[i][a] {
			return false
		}
	}
	for i, j := a, b; i > 0 && j > 0; i, j = i-1, j-1 {
		if sol[i][j] {
			return false
		}
	}
	for i, j := a, b; i < MAX_COL && j > 0; i, j = i+1, j-1 {
		if sol[i][j] {
			return false
		}
	}
	return true

}
func nQueen(sol [][]int, col int) bool {
	if col >= MAX_COL {
		return true
	}
	for i := 0; i < MAX_COL; i++ {
		if qSafe(sol, i, col) {
			sol[i][col] = 1
			if nQueen(sol, col+1) {
				return true
			}
			sol[i][col] = 0
		}

	}
	return false
}

func SolveNQueen(n int) bool {
	/*
		place one queen on the left most and bottom most position
		increment row by one || columan by one or both
		check
		    if there are any queens to the left in the same row
		    if there are any queens below or above
		    if there are any queens diagonally
		if yes then return false
		if no then place one queen and repeat until none are left.
	*/
}
