package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MIN_LINE_BUFFER_SIZE = 1024 * 1024      // 1 MB
const MAX_LINE_BUFFER_SIZE = 10 * 1024 * 1024 // 10 MB

func AskFile() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("CAMINHO DO FICHEIRO: ")
	path, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(path), nil
}

func GetFileData() (*bufio.Scanner, *bufio.Writer) {
	reader := bufio.NewScanner(os.Stdin)
	reader.Buffer(make([]byte, MIN_LINE_BUFFER_SIZE), MAX_LINE_BUFFER_SIZE)

	writer := bufio.NewWriter(os.Stdout)

	return reader, writer
}

func CheckError(err error) {
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}

func PrintResult(words []string) {
	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}
}
