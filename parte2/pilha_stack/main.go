package pilha_stack

import (
	"ProjetoAED/parte2/operations"
	"ProjetoAED/utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	var pilha operations.Pilha = &PilhaStack{}
	operations.StandardOperation(pilha, reader, writer)
}
