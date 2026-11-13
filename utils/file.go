package utils

import (
	"bufio"
	"strings"
	"unicode"
	"unicode/utf8"
)

func isLetter(r rune) bool {
	return unicode.IsLetter(r)
}

func isApostrophe(r rune) bool {
	return r == '\'' || r == '’'
}

func ScanWords(data []byte, atEOF bool) (int, []byte, error) {
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
			return start, nil, nil
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

	if i >= len(data) && !atEOF {
		return start, nil, nil
	}

	return i, data[start:i], nil
}

func StackWords(pilha Pilha, row string) {
	for {
		row = strings.TrimLeftFunc(row, unicode.IsSpace)
		if row == "" {
			return
		}
		index := strings.IndexFunc(row, unicode.IsSpace)
		if index == -1 {
			pilha.Push(row)
			return
		}
		pilha.Push(row[:index])
		row = row[index:]
	}
}

func WriteParagraph(pilha Pilha, writer *bufio.Writer) {
	first := true
	for !pilha.IsEmpty() {
		word, _ := pilha.Pop()
		if !first {
			writer.WriteString(" ")
		}
		writer.WriteString(word)
		first = false
	}
	writer.WriteString("\n")
}
