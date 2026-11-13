package main

import (
	"benchmarks_tests"
	"testing"
)

func BenchmarkSelectionSort(b *testing.B) {
	benchmarks_tests.BenchmarkN(b, "../1984.txt", selectionSort)
}
