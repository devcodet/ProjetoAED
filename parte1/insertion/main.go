package insertion

import (
	"ProjetoAED/utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	words, err := utils.ReadFile(reader, true, utils.ScanWords)
	utils.CheckError(err)

	insertionSort(words)

	utils.WriteFile(words, writer)
	utils.CheckError(err)
}
