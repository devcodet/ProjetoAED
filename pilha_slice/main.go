package main

import (
	"utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	var pilha utils.Pilha = &PilhaSlice{}
	utils.StandardOperation(pilha, reader, writer)
}
