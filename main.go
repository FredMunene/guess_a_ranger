package main

import (
	"fmt"
	"os"

	"tetris/src"
)

func main() {
	// handle no of arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [FILE_PATH]\n\nEX: go run . text1.txt")
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

	// resizing tetrimones
	for _, cube := range allTetrominoes {
		cube := src.ResizeTetri(cube)
		newTetrominoes = append(newTetrominoes, cube)
	}

	board := src.SolveBoard(newTetrominoes)
	for _, row := range board {
		fmt.Println(row)
	}
}
