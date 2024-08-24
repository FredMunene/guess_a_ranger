package src

import (
	"fmt"
	"os"
	"strings"
)

//ReadFile retrieves data from the path provided
func ReadFile(path string) ([]string, error) {
	fileContents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(fileContents) == 0{
		return nil, fmt.Errorf("ERROR")
	}

	buffer := string(fileContents)

	tetrominoes := strings.Split(buffer, "\n")

	return tetrominoes, nil
}
