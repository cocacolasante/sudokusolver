package main

import (
	"fmt"

	"github.com/cocacolasante/sudokusolver/algo"
	"github.com/cocacolasante/sudokusolver/board"
)


func main() {
	grid := board.CreateBoard()
	grid.FillRandomNumbers()
	fmt.Println(grid)

	algo.SolveSudoku(grid)
	for _, row := range grid {
		fmt.Println(row)
	}
}