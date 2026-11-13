package main

import (
	"benchmarks_tests"
	"testing"
)

func BenchmarkInsertionSort(b *testing.B) {
	benchmarks_tests.BenchmarkN(b, "../1984.txt", insertionSort)
}
