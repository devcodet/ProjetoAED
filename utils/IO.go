package utils

import (
	"bufio"
	"strings"
)

const DEFAULT_BUFFER_SIZE = 64 * 1024    // 64 KB
const MAX_BUFFER_SIZE = 10 * 1024 * 1024 // 10 MB

func ReadFile(scanner *bufio.Scanner, isLower bool, split bufio.SplitFunc) ([]string, error) {
	var words []string
	seen := make(map[string]struct{})

	scanner.Buffer(make([]byte, 0, DEFAULT_BUFFER_SIZE), MAX_BUFFER_SIZE)
	scanner.Split(split)

	for scanner.Scan() {
		w := scanner.Text()
		if isLower {
			w = strings.ToLower(w)
		}

		if _, exists := seen[w]; !exists {
			seen[w] = struct{}{}
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
