package pilha_slice

import (
	"ProjetoAED/parte2/operations"
	"ProjetoAED/utils"
	"strings"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	var pilha operations.Pilha = &PilhaSlice{}

	for reader.Scan() {
		row := reader.Text()

		if strings.TrimSpace(row) == "" {
			if !pilha.IsEmpty() {
				operations.WriteParagraph(pilha, writer)
			}
			writer.WriteString("\n")
		} else {
			operations.StackWords(row, pilha)
		}
	}
	utils.CheckError(reader.Err())

	if !pilha.IsEmpty() {
		operations.WriteParagraph(pilha, writer)
	}
}
