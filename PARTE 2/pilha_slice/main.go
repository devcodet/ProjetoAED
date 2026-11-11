package main_slice

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func (p *Pilha) Push(word string) {
	p.items = append(p.items, word)
}

func (p *Pilha) Pop() (string, bool) {
	if len(p.items) == 0 {
		return "", false
	}
	last := len(p.items) - 1
	word := p.items[last]
	p.items = p.items[:last]
	return word, true
}

func (p *Pilha) IsEmpty() bool {
	return len(p.items) == 0
}

func empilhawords(row string, p *Pilha) {
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

func escreveParagrafo(p *Pilha, out *bufio.Writer) {
	primeira := true
	for !p.IsEmpty() {
		w, _ := p.Pop()
		if !primeira {
			out.WriteString(" ")
		}
		out.WriteString(w)
		primeira = false
	}
	out.WriteString("\n")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "uso: go run main.go <ficheiro>")
		os.Exit(1)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "erro:", err)
		os.Exit(1)
	}
	defer f.Close()

	leitor := bufio.NewReader(f)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	p := &Pilha{}

	for {
		row, err := leitor.ReadString('\n')

		if strings.TrimSpace(row) == "" {
			if !p.IsEmpty() {
				escreveParagrafo(p, out)
			}
			if row != "" {
				out.WriteString("\n")
			}
		} else {
			empilhawords(row, p)
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, "erro a ler:", err)
			break
		}
	}

	if !p.IsEmpty() {
		escreveParagrafo(p, out)
	}
}
