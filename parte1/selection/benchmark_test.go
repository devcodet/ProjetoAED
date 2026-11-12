package selection

import (
	"parte1_benchmark_tests"
	"testing"
)

func BenchmarkSelectionSort(b *testing.B) {
	parte1_benchmark_tests.BenchmarkN(b, "../../1984.txt", selectionSort)
}
