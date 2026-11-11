package main

import (
	"bufio"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

const DEFAULT_BUFFER_SIZE = 64 * 1024    // 64 KB
const MAX_BUFFER_SIZE = 10 * 1024 * 1024 // 10 MB

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}

func isApostrophe(r rune) bool {
	return r == '\'' || r == '’'
}

func scanWords(data []byte, atEOF bool) (int, []byte, error) {
	start := 0

	for start < len(data) {
		r, w := utf8.DecodeRune(data[start:])
		if r == utf8.RuneError && w == 1 && !atEOF && !utf8.FullRune(data[start:]) {
			return start, nil, nil
		}

		if isLetter(r) {
			break
		}
		start += w
	}
	if start >= len(data) {
		if atEOF {
			return len(data), nil, nil
		}
		return start, nil, nil
	}

	i := start
	for i < len(data) {
		if !utf8.FullRune(data[i:]) && !atEOF {
			return start, nil, nil // pedir mais dados
		}
		r, w := utf8.DecodeRune(data[i:])

		if isLetter(r) {
			i += w
			continue
		}

		if isApostrophe(r) {
			next := i + w
			if next >= len(data) {
				if !atEOF {
					return start, nil, nil
				}
				break
			}
			if !utf8.FullRune(data[next:]) && !atEOF {
				return start, nil, nil
			}
			r2, _ := utf8.DecodeRune(data[next:])
			if isLetter(r2) {
				i = next
				continue
			}
		}
		break
	}

	// palavra pode continuar no próximo bloco
	if i >= len(data) && !atEOF {
		return start, nil, nil
	}

	return i, data[start:i], nil
}

func readFile(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var words []string
	seen := make(map[string]struct{})

	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, DEFAULT_BUFFER_SIZE), MAX_BUFFER_SIZE)
	scanner.Split(scanWords)

	for scanner.Scan() {
		w := strings.ToLower(scanner.Text())

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

func writeFile(path string, words []string) error {
	data := []byte(strings.Join(words, "\n"))
	if err := os.WriteFile(path, data, 0644); err != nil {
		checkError(err)
	}
	return nil
}
