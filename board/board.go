package board

import (
	"math/rand"
)

type SudokuGrid [9][9]int

func CreateBoard() *SudokuGrid {
	return &SudokuGrid{}
}

func(grid *SudokuGrid) FillRandomNumbers() {
	// Fill each row with random numbers
	for row := 0; row < 9; row++ {
		grid.FillRow(row)
	}
}

func(grid *SudokuGrid)FillRow(row int) {
	// Create a slice containing the numbers 1 to 9
	availableNumbers := rand.Perm(9) // Shuffle the numbers

	// Iterate through each column in the row
	for col := 0; col < 9; col++ {
		// Try placing a random number in the current cell
		for _, num := range availableNumbers {
			if IsValid(grid, row, col, num+1) {
				grid[row][col] = num + 1
				break
			}
		}
	}
}

func IsValid(grid *SudokuGrid, row, col, num int) bool {
	// checks the line
	for i := 0; i < 9; i++ {
		if grid[row][i] == num {
			return false
		}
	}

	// checks the column

	for i := 0; i < 9; i++ {
		if grid[i][col] == num {
			return false
		}
	}

	// check the box
	startRow, startCol := row-row%3, col-col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true

}
