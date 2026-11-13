package utils

import (
	"bufio"
	"os"
	"strings"
)

const DEFAULT_BUFFER_SIZE = 64 * 1024    // 64 KB
const MAX_BUFFER_SIZE = 10 * 1024 * 1024 // 10 MB

const MIN_LINE_BUFFER_SIZE = 1024 * 1024      // 1 MB
const MAX_LINE_BUFFER_SIZE = 10 * 1024 * 1024 // 10 MB

func ReadFile(scanner *bufio.Scanner, split bufio.SplitFunc, maxWords int) ([]string, error) {
	var words []string
	seen := make(map[string]struct{})

	scanner.Buffer(make([]byte, 0, DEFAULT_BUFFER_SIZE), MAX_BUFFER_SIZE)
	scanner.Split(split)

	for scanner.Scan() {
		word := scanner.Text()
		w := strings.ToLower(word)

		if _, exists := seen[w]; !exists {
			seen[w] = struct{}{}

			if maxWords > 0 && len(words) >= maxWords {
				break
			}
			words = append(words, w)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

func WriteFile(words []string, writer *bufio.Writer) error {
	data := []byte(strings.Join(words, "\n"))
	if _, err := writer.Write(data); err != nil {
		return err
	}
	return nil
}

func GetFileDataFromFile(filePath string) (*bufio.Scanner, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewScanner(file)
	reader.Buffer(make([]byte, MIN_LINE_BUFFER_SIZE), MAX_LINE_BUFFER_SIZE)

	return reader, nil
}

func GetFileData() (*bufio.Scanner, *bufio.Writer) {
	reader := bufio.NewScanner(os.Stdin)
	reader.Buffer(make([]byte, MIN_LINE_BUFFER_SIZE), MAX_LINE_BUFFER_SIZE)

	writer := bufio.NewWriter(os.Stdout)

	return reader, writer
}
