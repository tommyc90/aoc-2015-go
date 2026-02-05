package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ReadSingleLine(filePath string) (string, error) {
	lines, err := ReadLines(filePath)
	if err != nil {
		return "", err
	}

	if len(lines) != 1 {
		return "", fmt.Errorf("invalid input: expected 1 line, got %d", len(lines))
	}

	return lines[0], nil
}
