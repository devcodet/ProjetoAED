package utils

import (
	"ProjetoAED/utils"
	"bufio"
	"os"
	"strings"
)

const DEFAULT_BUFFER_SIZE = 64 * 1024    // 64 KB
const MAX_BUFFER_SIZE = 10 * 1024 * 1024 // 10 MB

func ReadFile(path string, isLower bool, split bufio.SplitFunc) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var words []string
	seen := make(map[string]struct{})

	scanner := bufio.NewScanner(f)
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

func WriteFile(path string, words []string) error {
	data := []byte(strings.Join(words, "\n"))
	if err := os.WriteFile(path, data, 0644); err != nil {
		utils.CheckError(err)
	}
	return nil
}
