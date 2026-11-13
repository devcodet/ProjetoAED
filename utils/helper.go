package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}

func PrintResult(words []string) {
	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}
}

func StandardOperation(pilha Pilha, reader *bufio.Scanner, writer *bufio.Writer) {
	for reader.Scan() {
		row := reader.Text()

		if strings.TrimSpace(row) == "" {
			if !pilha.IsEmpty() {
				WriteParagraph(pilha, writer)
			}
		} else {
			StackWords(pilha, row)
		}

		CheckError(reader.Err())
	}

	if !pilha.IsEmpty() {
		WriteParagraph(pilha, writer)
	}
}
