package main

import (
	"fmt"
	"os"
)

func is_valid(grid [][]int, x int, y int, num int) bool {
	// Check row
	for i := 0; i < 9; i++ {
		if grid[x][i] == num {
			return false
		}
	}
	// Check column
	for i := 0; i < 9; i++ {
		if grid[i][y] == num {
			return false
		}
	}
	// Check box
	box_x := (x / 3) * 3
	box_y := (y / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[box_x+i][box_y+j] == num {
				return false
			}
		}
	}
	return true
}

func solve(grid [][]int) bool {
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if grid[x][y] == 0 {
				for num := 1; num <= 9; num++ {
					if !is_valid(grid, x, y, num) {
						continue
					}
					grid[x][y] = num
					if solve(grid) {
						return true
					} else {
						grid[x][y] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func print_grid(grid [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
}

func main() {
	// Parse input arguments
	var grid [][]int
	for _, arg := range os.Args[1:] {
		row := make([]int, 9)
		for i, c := range arg {
			if c == '.' {
				row[i] = 0
			} else {
				row[i] = int(c - '0')
			}
		}
		grid = append(grid, row)
	}

	// Solve Sudoku
	if solve(grid) {
		print_grid(grid)
	} else {
		fmt.Println("No solution exists.")
	}
}
