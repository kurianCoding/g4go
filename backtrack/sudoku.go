package backtrack

/* solving the sudoku puzzle using backtracking*/
/* check if cell is empty or not*/
func emptyRowsCols(g [][]int) bool {
	/* range rows and columns*/
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if g[i][j] == 0 { /* found an empty cell*/
				return true
			}
		}
	}
	return false
}
func isColSafe(g [][]int, row, a int) bool {
	for i := 0; i < 9; i++ {
		if g[i][row] == a {
			return false
		}
	}
	return true
}
func isRowSafe(g [][]int, row, a int) bool {
	for i := 0; i < 9; i++ {
		if g[row][i] == a {
			return false
		}
	}
	return true
}
func isGridSafe(g [][]int, row, col, a int) bool {
	row = row / 3 /*getting the first row of div where cell belongs*/
	col = col / 3 /*geeting the first col of div where cell belongs*/
	for i := row * 3; i < 3*(row+1); i++ {
		for j := col * 3; j < 3*(col+1); j++ {
			if g[i][j] == a {
				return false
			}
		}
	}
	return true
}

func isSafeSudoKu(g [][]int, row, col, a int) bool {
	return isColSafe(g, col, a) && isRowSafe(g, row, a) && isGridSafe(g, row, col, a)
}

/*recurring function*/
func sudoUtil(grid [][]int) bool {
	/*check if there are empty rows or columnns*/
	if !emptyRowsCols(grid) {
		return true
	}

	/* if there is row or column empty, check if the number is safe to enter*/
	/* check if this number was used in row*/
	/* check if this number was used in col*/
	/*chekc if this number was used in 9x9 grid*/

	/* if the number is safe to enter, udpate the cell and find the next number*/
	return false
}
func SD(grid [][]int) bool {
	/* call the recurring function*/
	return sudoUtil(grid)
}
