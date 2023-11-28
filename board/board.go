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
	// Check if the number is not present in the current row and column
	for i := 0; i < 9; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}

	// Check if the number is not present in the current 3x3 subgrid
	startRow, startCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

