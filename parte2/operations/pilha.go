package operations

import (
	"bufio"
	"strings"
	"unicode"
)

type Pilha interface {
	Push(word string)
	Pop() (string, bool)
	IsEmpty() bool
	Empty() bool
}

func StackWords(pilha Pilha, row string) {
	for {
		row = strings.TrimLeftFunc(row, unicode.IsSpace)
		if row == "" {
			return
		}
		i := strings.IndexFunc(row, unicode.IsSpace)
		if i == -1 {
			pilha.Push(row)
			return
		}
		pilha.Push(row[:i])
		row = row[i:]
	}
}

func WriteParagraph(pilha Pilha, writer *bufio.Writer) {
	first := true
	for !pilha.IsEmpty() {
		w, _ := pilha.Pop()
		if !first {
			writer.WriteString(" ")
		}
		writer.WriteString(w)
		first = false
	}
	writer.WriteString("\n")
}
