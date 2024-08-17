package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, fmt.Errorf("failed to read line in file")
	}

	file.Close()
	return lines, nil
}
