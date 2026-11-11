package main

import (
	"ProjetoAED/parte1/file"
	"ProjetoAED/utils"
)

func main() {
	path, err := utils.AskFile()

	utils.CheckError(err)

	words, err := utils.ReadFile(path, true, file.ScanWords)
	utils.CheckError(err)

	selectionSort(words)
	utils.WriteFile("1_selection.txt", words)

	insertionSort(words)
	utils.WriteFile("1_insertion.txt", words)

	words = mergeSort(words)
	utils.WriteFile("1_mergesort.txt", words)

	utils.PrintResult(words)
}
