package src

import "errors"

const errMessage = "ERROR"
// ValidTetrimino checks a tetromino is valid
func ValidTetrimino(cube []string) (bool, error) {
	cellCount := 0
	rowCount := 0

	for _, row := range cube {
		rowCount++
		for _, cell := range row {
			if cell == '#' {
				cellCount++
			}
		}
	}
	if cellCount != 4 || rowCount != 4 {
		return false, errors.New(errMessage)
	}

	if !Connections(cube) {
		return false, errors.New(errMessage)
	}
	return true, nil
}

// Connections returns a bool if the tetromino has valid number of connections
func Connections(cube []string) bool {
	connect := 0
	for i, row := range cube {
		for j, cell := range row {
			if cell == '#' {
				if j < 3 {
					if row[j+1] == '#' {
						connect++
					}
				}
				if j > 0 {
					if row[j-1] == '#' && j > 0 {
						connect++
					}
				}
				if i > 0 {
					if cube[i-1][j] == '#' {
						connect++
					}
				}
				if i < 3 {
					if cube[i+1][j] == '#' { 
						connect++
					}
				}

			}
		}
	}
	if connect == 8 || connect == 6 {
		return true
	}
	return false
}
