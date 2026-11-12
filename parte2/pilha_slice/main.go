package pilha_slice

import (
	"ProjetoAED/parte2/operations"
	"ProjetoAED/utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	var pilha operations.Pilha = &PilhaSlice{}
	operations.StandardOperation(pilha, reader, writer)
}
