package backtrack

import (
	"fmt"
)

const MAX_X int = 8
const MAX_Y int = 8

func isValidMove(path *[][]int, x, y int) bool {
	if x > 0 && x < MAX_X-1 && y > 0 && y < MAX_Y-1 && (*path)[x][y] == -1 {
		return true
	}
	return false
}

func PrintValid(path [][]int) {
	for row := 0; row < MAX_X; row++ {
		for col := 0; col < MAX_Y; col++ {
			fmt.Printf("%d", path[row][col])
		}
		fmt.Printf("\n")
	}
	return
}
func knightsUtility(path *[][]int, xMove []int, yMove []int, x, y, moveCount int) bool {

	if moveCount == MAX_X*MAX_Y {
		return true
	}
	for i := 0; i < len(xMove); i++ {
		xDel := xMove[i]
		yDel := yMove[i]
		if isValidMove(path, x+xDel, y+yDel) {
			(*path)[x+xDel][y+yDel] = moveCount
			if knightsUtility(path, xMove, yMove, x+xDel, y+yDel, moveCount+1) {
				return true
			} else {
				(*path)[x+xDel][y+yDel] = -1
			}
		}
	}
	return false
}
func KnightsTour(x, y int) bool {
	path := [][]int{
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1, -1, -1},
	}
	xMove := []int{2, 1, -2, -1, 2, 1, -2, -1}
	yMove := []int{1, 2, 1, 2, -1, -2, 1, 2}
	path[x][y] = 0
	if knightsUtility(&path, xMove, yMove, x, y, 0) {
		PrintValid(path)
		return true
	} else {
		return false
	}
	return true
}
