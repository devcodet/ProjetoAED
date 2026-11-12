package operations

import (
	"bufio"
	"strings"
	"utils"
)

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

		utils.CheckError(reader.Err())
	}

	if !pilha.IsEmpty() {
		WriteParagraph(pilha, writer)
	}
}
