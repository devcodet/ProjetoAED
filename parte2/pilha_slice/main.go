package main

import (
	"ProjetoAED/utils"
	"bufio"
	"io"
	"os"
	"strings"
	"unicode"
)

func stackWords(row string, p *Pilha) {
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
	path, err := utils.AskFile()

	f, err := os.Open(path)
	utils.CheckError(err)

	defer f.Close()

	reader := bufio.NewReader(f)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	p := &Pilha{}

	for {
		row, err := reader.ReadString('\n')

		if strings.TrimSpace(row) == "" {
			if !p.IsEmpty() {
				writeParagraph(p, out)
			}
			if row != "" {
				out.WriteString("\n")
			}
		} else {
			stackWords(row, p)
		}

		if err == io.EOF {
			break
		}
		utils.CheckError(err)
	}

	if !p.IsEmpty() {
		writeParagraph(p, out)
	}
}
