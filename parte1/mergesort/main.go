package mergesort

import (
	"ProjetoAED/parte1/file"
	"ProjetoAED/utils"
)

func main() {
	reader, writer := utils.GetFileData()

	words, err := utils.ReadFile(reader, true, file.ScanWords)
	utils.CheckError(err)

	words = mergeSort(words)

	utils.WriteFile(words, writer)
	utils.CheckError(err)
}
