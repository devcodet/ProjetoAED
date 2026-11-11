package main_stack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func (p *Pilha) Push(word string) {
	p.topo = &Node{word: word, next: p.topo}
}

func (p *Pilha) Pop() (string, bool) {
	if p.topo == nil {
		return "", false
	}
	word := p.topo.word
	p.topo = p.topo.next
	return word, true
}

func (p *Pilha) IsEmpty() bool {
	return p.topo == nil
}

// empilhawords separa as words da row (sem usar slices) e faz Push de cada uma.
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

// writeParagraph esvazia a pilha, escrevendo as words numa só row.
func writeParagraph(p *Pilha, out *bufio.Writer) {
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
			// row em branco: fim de parágrafo
			if !p.IsEmpty() {
				writeParagraph(p, out)
				out.WriteString("\n") // mantém a separação entre parágrafos
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

	// último parágrafo (se o ficheiro não terminar com row em branco)
	if !p.IsEmpty() {
		writeParagraph(p, out)
	}
}
