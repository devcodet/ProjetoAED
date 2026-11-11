package main

func main() {
	path, err := askFile()

	checkError(err)

	words, err := readFile(path)
	checkError(err)

	printResult(words)
}
