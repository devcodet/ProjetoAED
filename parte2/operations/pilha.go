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
}

func StackWords(row string, p Pilha) {
	for {
		row = strings.TrimLeftFunc(row, unicode.IsSpace)
		if row == "" {
			return
		}
		i := strings.IndexFunc(row, unicode.IsSpace)
		if i == -1 {
			p.Push(row)
			return
		}
		p.Push(row[:i])
		row = row[i:]
	}
}

func WriteParagraph(p Pilha, out *bufio.Writer) {
	first := true
	for !p.IsEmpty() {
		w, _ := p.Pop()
		if !first {
			out.WriteString(" ")
		}
		out.WriteString(w)
		first = false
	}
	out.WriteString("\n")
}
