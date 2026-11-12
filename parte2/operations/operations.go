package operations

import (
	"ProjetoAED/utils"
	"bufio"
	"strings"
)

func StandardOperation(pilha Pilha, reader *bufio.Scanner, writer *bufio.Writer) {
	for reader.Scan() {
		row := reader.Text()

		if strings.TrimSpace(row) == "" {
			// row em branco: fim de parágrafo
			if !pilha.IsEmpty() {
				WriteParagraph(pilha, writer)
				writer.WriteString("\n") // mantém a separação entre parágrafos
			}
		} else {
			StackWords(row, pilha)
		}

		utils.CheckError(reader.Err())
	}

	if !pilha.IsEmpty() {
		WriteParagraph(pilha, writer)
	}
}
