package src

import (
	"os"
	"strings"
)

func ReadFile(path string) ([]string, error) {
	fileContents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	buffer := string(fileContents)

	tetrominoes := strings.Split(buffer, "\n")

	return tetrominoes, nil
}
