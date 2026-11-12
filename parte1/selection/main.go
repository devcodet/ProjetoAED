package selection

import (
	"utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	words, err := utils.ReadFile(reader, utils.ScanWords, 0)
	utils.CheckError(err)

	selectionSort(words)

	utils.WriteFile(words, writer)
	utils.CheckError(err)
}
