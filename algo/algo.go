package algo

import (
	"github.com/cocacolasante/sudokusolver/board"
)

func SolveSudoku(grid *board.SudokuGrid) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				for num := 1; num <= 9; num++ {
					if board.IsValid(grid, i, j, num) {
						grid[i][j] = num

						if SolveSudoku(grid) {
							return true
						} else {
							grid[i][j] = 0
						}
					}
				}
				return false
			}
		}
	}
	return true
}
