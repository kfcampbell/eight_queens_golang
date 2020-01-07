package main

import (
	"fmt"
	"log"
)

type Game struct {
	Board []([]bool)
}

func printBoard(game Game) {
	for _, row := range game.Board {
		for _, queenIsPlaced := range row {
			//log.Printf("Location: [%v, %v]. Queen placement: %v", i, j, queenIsPlaced)
			fmt.Printf(" %v ", transformBoolForPrinting(queenIsPlaced))
		}
		fmt.Printf("\n")
	}
}

func transformBoolForPrinting(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

func isSafeToPlaceQueen(game Game, row, col int) bool {

	// check the row, hold the column constant
	for j := 0; j < len(game.Board); j++ {
		if game.Board[row][j] == true {
			return false
		}
	}

	// check the column, hold the row constant
	for i := 0; i < len(game.Board[row]); i++ {
		if game.Board[i][col] == true {
			return false
		}
	}

	// check upper left diagonal
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if game.Board[i][j] == true {
			return false
		}
	}

	// check lower left diagonal
	for i, j := row, col; i < len(game.Board) && j >= 0; i, j = i+1, j-1 {
		if game.Board[i][j] == true {
			return false
		}
	}

	return true
}

func solveNQueens(game Game, col int) bool {
	// base case: if all queens are placed (we're at or greater than our total column count), return true
	if col >= len(game.Board) {
		return true
	}

	// freeze the current column and try placing the queen in all rows
	for i := 0; i < len(game.Board[0]); i++ {
		if isSafeToPlaceQueen(game, i, col) {
			// place the queen
			game.Board[i][col] = true

			// recurse to place the rest of the queens
			if solveNQueens(game, col+1) == true {
				return true
			}

			// otherwise remove the current queen and backtrack
			game.Board[i][col] = false
		}
	}

	// return false if we can't place the queen anywhere
	return false
}

func generateBoard(sizeLength int) []([]bool) {
	board := make([]([]bool), sizeLength)
	for i := 0; i < sizeLength; i++ {
		currSlice := make([]bool, sizeLength)
		board[i] = currSlice
	}
	return board
}

func main() {
	log.Printf("Program executing.")

	game := Game{
		Board: generateBoard(8),
	}

	if solveNQueens(game, 0) == false {
		log.Fatalf("Solution does not exist.")
	}

	printBoard(game)

	log.Printf("Program done executing.")
}
