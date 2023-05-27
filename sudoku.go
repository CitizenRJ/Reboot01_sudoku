package main

import (
	"fmt"
	"os"
)

// Husain Hanoon

func ValidateArgs() bool {
	args := os.Args[1:]
	// the operating system and arguements,
	if len(args) != 9 {
		return false
	}
	//it will check the number of rows from 1 to 9.
	for _, arg := range args {
		if len(arg) != 9 {
			return false
		}
		//the range of argument, it will check the number of boxes from 1 to 9 or (.) .
		for _, r := range arg {
			if (r >= '1' && r <= '9') || r == '.' {
			} else {
				return false
			}
		}
	}
	// the numbers in rows should be from 1 to 9.
	return true
}

// Fatima
func NoDuplicate(board [][]int, row, col, num int) bool {
	// check if the number is already in the row
	for i := 0; i < 9; i++ {
		if board[row][i] == num && i != col {
			return false
		}
	}

	// check if the number is already in the column
	for i := 0; i < 9; i++ {
		if board[i][col] == num && row != i {
			return false
		}
	}

	// check if the number is already in the 3x3 grid // i for the box inside the row / j box in culomn
	gridRow, gridCol := row/3*3, col/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[gridRow+i][gridCol+j] == num && (gridRow+i != row && gridCol+j != col) {
				return false
			}
		}
	}

	return true
}

// Mina
func main() {
	if !ValidateArgs() {
		fmt.Println("Error")
		return
	}
	// to build the table
	table := BuildTable()
	// check duplicates
	for x, row := range table {
		for y, n := range row {
			if table[x][y] == 0 {
				continue
			}
			if !NoDuplicate(table, x, y, n) {
				fmt.Println("Error")
				return
			}
		}
	}
	solveSudoku(table)
	printBoard(table)

}

// Mina
func solveSudoku(board [][]int) bool {
	row, col := 0, 0
	foundEmptyCell := false

	// find the next empty cell
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				row, col = i, j
				foundEmptyCell = true
				break
			}
		}
		if foundEmptyCell {
			break
		}
	}

	// if there are no more empty cells, the Sudoku is solved
	if !foundEmptyCell {
		return true
	}

	// try to fill in the empty cell with a valid number
	for num := 1; num <= 9; num++ {
		if isValid(board, row, col, num) {
			board[row][col] = num
			if solveSudoku(board) {
				return true
			}
			board[row][col] = 0
		}
	}

	// if no valid number was found, backtrack
	return false
}

// Mina
func isValid(board [][]int, row, col, num int) bool {
	// check if the number is already in the row
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// check if the number is already in the column
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// check if the number is already in the 3x3 grid
	gridRow, gridCol := row/3*3, col/3*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[gridRow+i][gridCol+j] == num {
				return false
			}
		}
	}

	return true
}

// Hussian Hanoon
func printBoard(board [][]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(board[i][j])
			if j != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// Fatima
func BuildTable() [][]int {
	args := os.Args[1:]

	_table := make([][]int, 9)

	for i := 0; i < 9; i++ {
		_table[i] = make([]int, 9)
	}
	for x, row := range _table {
		for y := range row {
			if args[x][y] == '.' {
				_table[x][y] = 0
			} else {
				_table[x][y] = int(args[x][y] - 48)
			}

		}
	}
	return _table
}
