package pilha_list

import (
	"operations"
	"utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	var pilha operations.Pilha = &PilhaList{}
	operations.StandardOperation(pilha, reader, writer)
}
