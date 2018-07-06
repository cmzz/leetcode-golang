package main

import (
	"strconv"
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	rows := [9][9]int{}
	cols := [9][9]int{}
	cells := [9][9]int{}

	// i1 行号 i2 列号
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] >= '1' && board[i][j] <= '9' {
				cs := string(board[i][j])
				c, _ := strconv.Atoi(cs)
				c = c-1

				if rows[i][c] == 1 || cols[j][c] == 1 || cells[3 * (i/3) + j/3][c] == 1 {
					return false
				}

				rows[i][c] = 1
				cols[j][c] = 1
				cells[3 * (i/3) + j/3][c] = 1
			}
		}
	}

	return true
}

func main() {
	board := [][]byte {{'.','.','4','.','.','.','6','3','.'},{'.','.','.','.','.','.','.','.','.'},{'5','.','.','.','.','.','.','9','.'},{'.','.','.','5','6','.','.','.','.'},{'4','.','3','.','.','.','.','.','1'},{'.','.','.','7','.','.','.','.','.'},{'.','.','.','5','.','.','.','.','.'},{'.','.','.','.','.','.','.','.','.'},{'.','.','.','.','.','.','.','.','.'}}
	fmt.Println(isValidSudoku(board))
}
