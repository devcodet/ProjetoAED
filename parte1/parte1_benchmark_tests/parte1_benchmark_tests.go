package parte1_benchmark_tests

import (
	"fmt"
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
