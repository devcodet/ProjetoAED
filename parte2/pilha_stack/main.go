package pilha_stack

import (
	"ProjetoAED/parte2/operations"
	"ProjetoAED/utils"
	"strings"
)

func main() {

	reader, writer := utils.GetFileData()
	defer writer.Flush()

	var pilha operations.Pilha = &PilhaStack{}

	for reader.Scan() {
		row := reader.Text()

		if strings.TrimSpace(row) == "" {
			// row em branco: fim de parágrafo
			if !pilha.IsEmpty() {
				operations.WriteParagraph(pilha, writer)
				writer.WriteString("\n") // mantém a separação entre parágrafos
			}
		} else {
			operations.StackWords(row, pilha)
		}

		utils.CheckError(reader.Err())
	}

	if !pilha.IsEmpty() {
		operations.WriteParagraph(pilha, writer)
	}
}
