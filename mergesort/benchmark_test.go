package main

import (
	"benchmarks_tests"
	"testing"
)

func BenchmarkMergeSort(b *testing.B) {
	benchmarks_tests.BenchmarkN(b, "../1984.txt", func(w []string) {
		_ = mergeSort(w)
	})
}
