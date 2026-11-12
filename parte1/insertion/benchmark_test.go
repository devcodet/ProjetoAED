package insertion

import (
	"parte1_benchmark_tests"
	"testing"
)

func BenchmarkInsertionSort(b *testing.B) {
	parte1_benchmark_tests.BenchmarkN(b, "../../1984.txt", insertionSort)
}
