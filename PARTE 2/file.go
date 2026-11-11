package utils

import (
	"os"
	"strings"
)

func writeFile(path string, words []string) error {
	data := []byte(strings.Join(words, "\n"))
	if err := os.WriteFile(path, data, 0644); err != nil {
		checkError(err)
	}
	return nil
}
