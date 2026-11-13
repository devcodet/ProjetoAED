package main

import (
	"utils"
)

func main() {
	reader, writer := utils.GetFileData()
	defer writer.Flush()

	words, err := utils.ReadFile(reader, utils.ScanWords)
	utils.CheckError(err)

	words = mergeSort(words)

	utils.WriteFile(words, writer)
	utils.CheckError(err)
}
