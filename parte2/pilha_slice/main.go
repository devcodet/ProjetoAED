package pilha_slice

import (
	"operations"
	"utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	var pilha operations.Pilha = &PilhaSlice{}
	operations.StandardOperation(pilha, reader, writer)
}
