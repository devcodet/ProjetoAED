package pilha_list

import (
	"utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	var pilha utils.Pilha = &PilhaList{}
	utils.StandardOperation(pilha, reader, writer)
}
