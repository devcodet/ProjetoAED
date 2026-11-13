package main

import (
	"benchmarks_tests"
	"testing"
)

func BenchmarkPilhaSlicePushPop(b *testing.B) {
	benchmarks_tests.BenchmarkPilhaPushPop(b, &PilhaSlice{}, "../1984.txt")
}
