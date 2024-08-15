package main

import (
	"fmt"
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

	for _, cube := range allTetrominoes {
	cube := src.ResizeTetri(cube)
		for _, line := range cube {
			println(line)
		}
		println("*****")
	}
	// TODO
	// CREATE SUPER BOARD
	// PLACING TETROMINOS ON BOARD
	// RECURSION
	// FINAL BOARD
}

