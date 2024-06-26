package helpers

import (
	"bufio"
	"os"
)

func ReadLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
func AppendToFile(filePath, content string) error {
	// Open the file in append mode, or create it if it doesn't exist
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append the content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
