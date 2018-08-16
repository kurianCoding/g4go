package backtrack

/*given a matrix of  0s and 1s, a rat is required to navigate
from the top left 0,0 to bottom right N-1,N-1.

    the function given here takes in the initial point and the
maze as input. It gives back as output a boolean that says if
the path exists of not
*/

/* write a function to cycle thrrough all possible movements
   to right and to down.
   if final row and final column is reached, it should, return
   true
   if no valid movement exist and final column is not reached
   it should return false

   there should be a function to check if the move is valid
*/
func isSafe(maze [][]int, x, y, x_max, y_max int) bool {
	if x >= 0 && y >= 0 && x <= x_max && y <= y_max && maze[x][y] == 1 {
		return true
	}
	return false
}
func scanMazeUtil(maze [][]int, x, y, x_dest, y_dest int) bool {
	if x == x_dest && y == y_dest {
		return true // search reached destination
	}
	if isSafe(maze, x, y, x_dest, y_dest) {
		if scanMazeUtil(maze, x+1, y, x_dest, y_dest) {
			return true
		}
		if scanMazeUtil(maze, x, y+1, x_dest, y_dest) {
			return true
		}
		return false
	}
	return false
}
func scanMaze(maze [][]int, x, y int) bool {
	x_dest := len(maze[0]) - 1
	y_dest := len(maze[0]) - 1
	return scanMazeUtil(maze, x, y, x_dest, y_dest)
}

func RatInMaze(maze [][]int, x, y int) bool {
	return scanMaze(maze, 0, 0)
}
