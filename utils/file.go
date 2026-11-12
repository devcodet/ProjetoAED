package utils

import (
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

	if i >= len(data) && !atEOF {
		return start, nil, nil
	}

	return i, data[start:i], nil
}
