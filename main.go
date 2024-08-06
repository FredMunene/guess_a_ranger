package main

import (
	"fmt"
	"os"

	"tetris/src"
)

func main() {
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

	allTetrominoes,err := src.GetTetromino(tetrimino)
	if err != nil{
		println(err)
		return
	}

	println(len(allTetrominoes))
	for i, cube := range allTetrominoes {
		// println("aaaa")
		for _, line := range cube {
			println(i, line)
		}
	}
}
