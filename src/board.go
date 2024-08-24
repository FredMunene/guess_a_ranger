package src

import "math"

// SolveBoard attempts to place all the given tetrominos on the smallest possible square board.
// It starts with an initial board size and increases the size if necessary until a valid
// placement is found for all tetrominos.
// 
// Parameters:
// - tetro: A 2D slice of strings where each sub-slice represents a tetromino.
//
// Returns:
// - A 1D slice of strings representing the board with all tetrominos placed,
//   or the smallest board that can accommodate all tetrominos.
//
// The function continuously tries to solve the problem by calling the solveRecurs
// function. If a solution is not found, the board size is increased and the function
// retries until a valid solution is found.

func SolveBoard(tetros [][]string) []string {
	InitialHeight := math.Ceil(math.Sqrt(float64(len(tetros) * 4)))
	board := CreateBoard(int(InitialHeight))

	for {
		if solveRecurs(tetros, board, 0) {
			return board
		}

		InitialHeight++
		board = CreateBoard(int(InitialHeight))

	}
}

// solveRecurs is a recursive function that attempts to place tetrominos on the board
// starting from a given index.
//
// Parameters:
// - tetro: A 2D slice of strings where each sub-slice represents a tetromino.
// - board: A 1D slice of strings representing the current state of the board.
// - index: The current index of the tetromino to be placed.
//
// Returns:
// - A boolean value: `true` if all tetrominos have been successfully placed on the board,
//   `false` otherwise.
//
// The function tries to place the current tetromino at all possible positions on the
// board. If a valid position is found, it recursively attempts to place the next tetromino.
// If no valid position is found for the current tetromino, the function backtracks by
// removing the last placed tetromino and tries the next possible position. If all positions
// have been exhausted without success, it returns `false` to indicate failure.

func solveRecurs(tetro [][]string, board []string, index int) bool {
	if index == len(tetro) {
		return true
	}

	tetromino := tetro[index]
	tetromino = letterise(tetromino, index)

	for i := range board {
		for j := range board[i] {
			if canPlace(tetromino, board, i, j) {

				place(tetromino, board, i, j)
				// trying to solve with next tetromino
				if solveRecurs(tetro, board, index+1) {
					// if the call suceeds, we've found a solution
					return true
				}
				// if the recursive call fails, remove tetromino and try the next position
				remove(tetromino, board, i, j) // backtracking
			}
		}
	}
	return false
}

// canPLace attempts to place tetrominos on the board starting from a given index.
// Parameters:
// - tetro: A 1D slice of strings representing a tetromino.
// - board: A 1D slice of strings representing the current state of the board.
// - index: The startingcurrent index of the board where the tetromino can be placed.
//
// Returns:
// - A boolean value: `true` if the tetromino can be successfully placed on the board,
//   `false` otherwise.
func canPlace(tetro, board []string, i, j int) bool {
	// check tetro is not out of bounds
	if i+len(tetro) > len(board) || j+len(tetro[0]) > len(board[0]) {
		return false
	}
	count := 0

	for r := 0; r < len(tetro); r++ {
		for c := 0; c < len(tetro[r]); c++ {
			// initiate board positions
			boardRow, boardCol := i+r, j+c

			if tetro[r][c] != '.' {
				if board[boardRow][boardCol] == '.' {
					count++
				} else {
					return false
				}
			}

		}
	}

	return count == 4
}

// pLace place tetrominos on the board starting from a given index.
// Parameters:
// - tetro: A 1D slice of strings representing a tetromino.
// - board: A 1D slice of strings representing the current state of the board.
// - index: The starting index of the board where the tetromino should be placed.
func place(tetro, board []string, i, j int) {
	for r := 0; r < len(tetro); r++ {
		for c := 0; c < len(tetro[r]); c++ {
			// corresponding position on the board
			boardRow, boardCol := i+r, j+c
			if tetro[r][c] != '.' {
				board[boardRow] = board[boardRow][0:boardCol] + string(tetro[r][c]) + board[boardRow][boardCol+1:]
			}
		}
	}
}

// remove replaces tetrominos on the board starting from a given index.
// Parameters:
// - tetro: A 1D slice of strings representing a tetromino.
// - board: A 1D slice of strings representing the current state of the board.
// - index: The starting index of the board where the tetromino should be removed.
func remove(tetro, board []string, i, j int) {
	for r := 0; r < len(tetro); r++ {
		for c := 0; c < len(tetro[r]); c++ {
			boardRow, boardCol := i+r, j+c
			if tetro[r][c] != '.' {
				board[boardRow] = board[boardRow][0:boardCol] + "." + board[boardRow][boardCol+1:]
			}
		}
	}
}
// letterise replaces `#` in the tetromino with a specific letter (A-Z)
// based on the provided index and converts other characters to `.`.
func letterise(tetro []string, letter int) []string {
	letter = letter + 65
	tetromino := []string{}
	for i := range tetro {
		line := ""
		for j := range tetro[i] {
			if tetro[i][j] == '#' {
				line += string(rune(letter))
			} else {
				line += "."
			}
		}
		tetromino = append(tetromino, line)
	}
	return tetromino
}
