package benchmarks_tests

import (
	"bufio"
	"fmt"
	"io"
	"testing"
	"utils"
)

func BenchmarkN(b *testing.B, inputFile string, sortFunc func([]string)) {
	reader, err := utils.GetFileDataFromFile(inputFile)
	utils.CheckError(err)

	words, err := utils.ReadFile(reader, utils.ScanWords)
	tests := []int{10, 100, 1000, 10000, 0}
	for _, n := range tests {
		size := n
		if size == 0 || size > len(words) {
			size = len(words)
		}

		new_words := make([]string, n)
		copy(new_words, words[:size])
		utils.CheckError(err)

		b.ResetTimer()

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.StartTimer()
			sortFunc(new_words)
			b.StopTimer()
			fmt.Printf("N=%d, Time=%v\n", n, b.Elapsed())
		})
	}
}

func BenchmarkPilhaPushPop(b *testing.B, pilha utils.Pilha, inputFile string) {
	reader, err := utils.GetFileDataFromFile(inputFile)
	utils.CheckError(err)

	words, err := utils.ReadFile(reader, utils.ScanWords)
	tests := []int{10, 100, 1000, 10000, 100000, 0}
	for _, n := range tests {
		size := n
		if size == 0 || size > len(words) {
			size = len(words)
		}

		new_words := make([]string, n)
		copy(new_words, words[:size])

		utils.CheckError(err)

		b.ReportAllocs()
		b.ResetTimer()

		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			b.StartTimer()
			for _, word := range new_words {
				pilha.Push(word)
			}
			for !pilha.IsEmpty() {
				pilha.Pop()
			}
			b.StopTimer()
			fmt.Printf("N=%d, Time=%v\n", n, b.Elapsed())
		})
	}
}

func BenchmarkPilhaStandardOperation(b *testing.B, pilha utils.Pilha, inputFile string) {
	reader, err := utils.GetFileDataFromFile(inputFile)

	writer := bufio.NewWriter(io.Discard)
	defer writer.Flush()

	utils.CheckError(err)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pilha.Empty()
		utils.StandardOperation(pilha, reader, writer)
	}
	b.StopTimer()
	fmt.Printf("N=%d, Time=%v\n", b.N, b.Elapsed())
}
