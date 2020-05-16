package main

import (
	"fmt"
	"time"
)

func main() {
	sudoku := [9][9]int{
		{7, 8, 0, 4, 0, 0, 1, 2, 0},
		{6, 0, 0, 0, 7, 5, 0, 0, 9},
		{0, 0, 0, 6, 0, 1, 0, 7, 8},
		{0, 0, 7, 0, 4, 0, 2, 6, 0},
		{0, 0, 1, 0, 5, 0, 9, 3, 0},
		{9, 0, 4, 0, 6, 0, 0, 0, 5},
		{0, 7, 0, 3, 0, 0, 0, 1, 2},
		{1, 2, 0, 0, 0, 7, 4, 0, 0},
		{0, 4, 9, 2, 0, 6, 0, 0, 7}}

	printSudoku(sudoku)
	start := time.Now()

	if solve(&sudoku) {
		fmt.Println("Sudoku solved !\nIt took ", time.Now().Sub(start).Microseconds(), " micro seconds")
		printSudoku(sudoku)
	} else {
		fmt.Print("Sudoku cannot be solved")
	}
}


// Print the sudoku
func printSudoku(sudoku [9][9]int) {
	fmt.Println("+———————+———————+———————+")
	for row:=0; row<9; row++ {
		fmt.Print("| ")
		for column:=0; column<9; column++ {
			fmt.Printf("%d ", sudoku[row][column])
			if column == 2 || column == 5 {
				fmt.Print("| ")
			}
		}
		fmt.Println("|")
		if row == 2 || row == 5 {
			fmt.Println("+———————+———————+———————+")
		}
	}
	fmt.Println("+———————+———————+———————+")
}


// Solve the sudoku recursively
func solve(sudoku *[9][9]int) bool {
	if !hasEmptyCell(sudoku) { // Meaning the sudoku is solved
		return true
	}
	for row:=0; row<9; row++ {
		for col :=0; col <9; col++ {
			if sudoku[row][col] == 0 { // Parsing each empty cell of the sudoku

				for candidate := 9; candidate >= 1; candidate-- {
					sudoku[row][col] = candidate
					if isSudokuValid(sudoku) {
						if solve(sudoku) {
							return true
						}
						sudoku[row][col] = 0
					} else {
						sudoku[row][col] = 0
					}
				}
				return false
			}
		}
	}
	return false
}


// Check if the sudoku contains cell with 0 values (empty)
// Returns true at least one value is equal to 0
func hasEmptyCell(sudoku *[9][9]int) bool {
	for row:=0; row<9; row++ {
		for column:=0; column<9; column++ {
			if sudoku[row][column] == 0 {
				return true
			}
		}
	}
	return false
}


// Check if the sudoku is valid according to the rules
// Return false if it detect an error
func isSudokuValid(sudoku *[9][9]int) bool {
	// Check each row for duplicates
	for row:=0; row<9; row++ {
		counter := [10]int{}
		for column:=0; column<9; column++ {
			counter[sudoku[row][column]]++
			if sudoku[row][column] != 0 && counter[sudoku[row][column]] > 1 {
				return false
			}
		}
	}

	// Check each column for duplicates
	for column:=0; column<9; column++ {
		counter := [10]int{}
		for row:=0; row<9; row++ {
			counter[sudoku[row][column]]++
			if sudoku[row][column] != 0 && counter[sudoku[row][column]] > 1 {
				return false
			}
		}
	}

	// Check each 3*3 section for duplicates
	for i:=0; i<9; i+=3 {
		for j:=0; j<9; j+=3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for column := j; column < j+3; column++ {
					counter[sudoku[row][column]]++
					if sudoku[row][column] != 0 && counter[sudoku[row][column]] > 1 {
						return false
					}
				}
			}
		}
	}

	return true
}

