package main

import (
	"benchmarks_tests"
	"testing"
)

func BenchmarkPilhaListPushPop(b *testing.B) {
	benchmarks_tests.BenchmarkPilhaPushPop(b, &PilhaList{}, "../1984.txt")
}

func BenchmarkPilhaListStandardOperation(b *testing.B) {
	benchmarks_tests.BenchmarkPilhaStandardOperation(b, &PilhaList{}, "../1984.txt")
}
