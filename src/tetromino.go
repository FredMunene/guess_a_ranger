package src

// ResizeTetri returns a slice of string with reduced height and width if row or column are blank
func ResizeTetri(cube []string) []string {
	newCube := []string{}

	// eliminate empty rows
	for i, row := range cube {
		if !isRowEmpty(row) {
			newCube = append(newCube, cube[i])
		}
	}

	// eliminate empty columns
	for {

		cols := 0
		if len(newCube) > 0 {
			cols = len(newCube[0])
		}

		cubeAltered := false
		resizedCube := []string{}

		for j := 0; j < cols; j++ {
			if isColumnEmpty(newCube, j) {
				for _, row := range newCube {
					if j != len(newCube)-1 {
						resizedCube = append(resizedCube, row[:j]+row[j+1:])
					} else {
						resizedCube = append(resizedCube, row[:j])
					}
				}
				newCube = resizedCube
				cubeAltered = true
				break
			}
		}

		if !cubeAltered {
			break
		}
	}

	// for _, line := range newCube {
	// 	println(line)
	// }
	// println(len(newCube))

	return newCube
}

// IsRowEmpty checks whether row is empty
func isRowEmpty(row string) bool {
	for _, char := range row {
		if char != '.' {
			return false
		}
	}
	return true
}

// IsColumnEmpty Checks whether column is empty
func isColumnEmpty(cube []string, colNum int) bool {
	col := ""
	for row := 0; row < len(cube); row++ {
		col += (string(cube[row][colNum]))
	}
	// println(col)
	return isRowEmpty(col)
}
