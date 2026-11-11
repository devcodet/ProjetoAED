package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func askFile() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("CAMINHO DO FICHEIRO: ")
	caminho, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(caminho), nil
}

func askSort() (typeSort, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("TIPO DE ORDENAÇÃO (1 - SELECTIONSORT, 2 - INSERTIONSORT, 3 - MERGESORT): ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return UnknownSort, err
	}

	n, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		return UnknownSort, fmt.Errorf("opção inválida: %q", strings.TrimSpace(input))
	}

	t := typeSort(n)
	switch t {
	case SelectionSort, InsertionSort, MergeSort:
		return t, nil
	default:
		return UnknownSort, fmt.Errorf("opção fora do intervalo: %d", n)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}

func printResult(words []string) {
	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}
}
