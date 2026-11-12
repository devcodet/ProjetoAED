package selection

import (
	"ProjetoAED/parte1/file"
	"ProjetoAED/utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	words, err := utils.ReadFile(reader, true, file.ScanWords)
	utils.CheckError(err)

	selectionSort(words)

	utils.WriteFile(words, writer)
	utils.CheckError(err)
}
