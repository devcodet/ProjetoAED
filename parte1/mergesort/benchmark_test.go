package mergesort

import (
	"parte1_benchmark_tests"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	parte1_benchmark_tests.BenchmarkN(b, "../../1984.txt", func(w []string) {
		_ = mergeSort(w)
	})
}
