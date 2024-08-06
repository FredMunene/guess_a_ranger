package src

import "fmt"

func GetTetromino(tetrimino []string)([][]string,error){
	allTetrimino := [][]string{}
	cube := []string{}
	for i := 0; i < len(tetrimino)-1; i++ {
		if tetrimino[i] != "" {
			cube = append(cube, tetrimino[i])
		} else {
			ok, _ := ValidTetrimino(cube)
			if !ok {
				//println("ERROR")
				return nil,fmt.Errorf(errMessage)
			}
			allTetrimino = append(allTetrimino, cube)
			cube = []string{}
		}
	}
	return allTetrimino,nil
}