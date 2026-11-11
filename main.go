package main

import "fmt"

func main() {
	path, err := askFile()
	typeSort, err := askSort()

	checkError(err)

	words, err := readFile(path)
	checkError(err)

	switch typeSort {
	case SelectionSort:
		selectionSort(words)
	case InsertionSort:
		insertionSort(words)
	case MergeSort:
		words = mergeSort(words)
	default:
		checkError(fmt.Errorf("TIPO DE ORDENAÇÃO INVÁLIDO: %d", int(typeSort)))
	}

	printResult(words)
}
