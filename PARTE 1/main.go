package main

func main() {
	path, err := askFile()

	checkError(err)

	words, err := readFile(path)
	checkError(err)

	selectionSort(words)
	writeFile("1_selection.txt", words)

	insertionSort(words)
	writeFile("1_insertion.txt", words)

	words = mergeSort(words)
	writeFile("1_mergesort.txt", words)

	printResult(words)
}
