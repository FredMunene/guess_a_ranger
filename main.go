package main

import (
	"fmt"
	"math"
	"os"

	"tetris/src"
)

func main() {
	// handle no of arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [FILE_PATH]\n\nEX: go run . templates/text1.txt")
		return
	}
	filePath := os.Args[1]

	tetrimino, err := src.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	// retrieve tetrominoes
	allTetrominoes, err := src.GetTetromino(tetrimino)
	if err != nil {
		fmt.Println(err)
		return
	}

	var newTetrominoes [][]string

	for _, cube := range allTetrominoes {
		cube := src.ResizeTetri(cube)

		newTetrominoes = append(newTetrominoes, cube)
		// for _, line := range cube {
		// 	println(line)
		// }
		// println("*****")
	}

	// TODO
	// CREATE SUPER BOARD
	InitialHeight := math.Ceil(math.Sqrt(float64(len(allTetrominoes) * 4)))
	board := src.CreateBoard(int(InitialHeight))
	// println(len(board))

	// PLACING TETROMINOS ON BOARD
	// RECURSION
	// FINAL BOARD
	Solver(board, newTetrominoes)
}

func Solver(board []string, allTetromino [][]string) {
	// alphabets A -> Z
	letter := 65
	for _, tetro := range allTetromino {
		if CanPlace(board, tetro) {
			Place(board, letterise(tetro, letter))
		}
		letter++
	}

	for _, line := range board {
		println(line)
	}
}

func CanPlace(board []string, tetro []string) bool {
	count := 0
	for i := range len(board) {
		for j := range len(board) {
			if board[i][j] == '.' && tetro[i][j] == '#' {
				count++
			}
		}
	}
	return count == 4
}

func Place(board []string, tetro []string) []string {
	for i := range len(board) {
		board[i], tetro[i] = tetro[i], board[i]
	}

	return board
}

func letterise(tetro []string, letter int) []string {
	tetromino := []string{}
	for i := range len(tetro) {
		line := ""
		for j := range len(tetro) {
			if tetro[i][j] == '#' {
				line += string(letter)
			}
		}
		tetromino = append(tetromino, line)
	}
	return tetromino
}
