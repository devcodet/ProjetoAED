package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskFile() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("CAMINHO DO FICHEIRO: ")
	caminho, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(caminho), nil
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
