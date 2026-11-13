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

	tests := []int{10, 100, 1000, 10000, 100000, 1000000, 10000000, 0}
	testing.Benchmark(func(b *testing.B) {
		for _, n := range tests {
			words, err := utils.ReadFile(reader, utils.ScanWords, n)
			utils.CheckError(err)

			b.ResetTimer()

			b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
				b.StartTimer()
				sortFunc(words)
				b.StopTimer()
				fmt.Printf("N=%d, Time=%v\n", n, b.Elapsed())
			})
		}
	})
}

func BenchmarkPilhaPushPop(b *testing.B, pilha utils.Pilha, inputFile string) {
	reader, err := utils.GetFileDataFromFile(inputFile)
	utils.CheckError(err)

	tests := []int{10, 100, 1000, 10000, 100000, 1000000, 10000000, 0}
	testing.Benchmark(func(b *testing.B) {
		for _, n := range tests {
			words, err := utils.ReadFile(reader, utils.ScanWords, n)
			utils.CheckError(err)

			b.ReportAllocs()
			b.ResetTimer()

			b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
				b.StartTimer()
				for _, word := range words {
					pilha.Push(word)
				}
				for !pilha.IsEmpty() {
					pilha.Pop()
				}
				b.StopTimer()
				fmt.Printf("N=%d, Time=%v\n", n, b.Elapsed())
			})
		}
	})
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
