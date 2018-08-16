package backtrack

/*
   this is a variation of nQueen problem. Here we are
   given a paritally filled chess board. The challenge is
   to find if the player can place one more queen on the
   chess board, without attacking existing queens

   given function takes a two dimensional map, this map contains
   1 for a queen and a 0 where no queen exists.

   It returns x,y as the location of the next queen that can be placed

   If no more queens can be placed -1,-1 is returned
*/

func fillRow(board *[][]int, row int) {
	for i := 0; i < row; i++ {
		(*board)[row][i] = 1
	}
	return
}

func fillCol(board *[][]int, col int) {

	for i := 0; i < row; i++ {
		(*board)[i][col] = 1
	}
	return
}

func fillDiag(board *[][]int, row, col int) {
	for i, j := row, col; i > 0 && j > 0; i, j = i-1, j-1 {
		(*board)[i][j] = 1
	}
	for i, j := row, col; i < MAX_COL && j < MAX_COL; i, j = i+1, j+1 {
		(*board)[i][j] = 1
	}
	return
}

func nFill(board *[][]int) {
	for i := 0; i < MAX_COL; i++ {
		for j := 0; i < MAX_COL; j++ {
			if (*board)[i][j] == 1 {
				fillRow(board, i)
				fillColl(board, j)
				fillDiag(board, i, j)
			}
		}
	}
}

func search(board *[][]int) (int, int) {
	var x, y int
	for i, j := 0, 0; i < MAX_COL && j < MAX_COL; i, j = i+1, j+1 {
		if (*board)[i][j] == 0 {
			x = i
			y = j
		}
	}
	return x, y
}

func nQueenVar1(board [][]int) (int, int) {
	/*
	   scan all the rows and columns of the board.
	   fill all the diagonal and crosswise on the board.

	   return the first unfilled cell
	*/
	nFill(&board)
	x, y := search(&board)
	return x, y
}
