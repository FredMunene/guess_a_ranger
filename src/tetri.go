package src

import "fmt"

// GetTetromino retrieves tetrominos from a slice of string
func GetTetromino(tetrimino []string) ([][]string, error) {
	allTetrimino := [][]string{}
	// if (len(tetrimino)-1 )% 5 != 0 {
	// 	fmt.Print("not ok")
	// 	println(len(tetrimino))
	// }
	cube := []string{}
	for i := 0; i < len(tetrimino)-1; i++ {
		
		if tetrimino[i] != "" {
			if len(tetrimino[i]) != 4 {
				return nil, fmt.Errorf("ERROR")
			}
			cube = append(cube, tetrimino[i])
		} else {
			ok, _ := ValidTetrimino(cube)
			if !ok {
				return nil, fmt.Errorf(errMessage)
			}
			allTetrimino = append(allTetrimino, cube)
			cube = []string{}
		}
	}
	return allTetrimino, nil
}

func CreateBoard(height int) []string {
	board := []string{}

	for i := 0; i < height; i++ {
		line := ""
		for j := 0; j < height; j++ {
			line += "."
		}
		board = append(board, line)
	}
	return board
}
