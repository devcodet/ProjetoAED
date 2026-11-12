package pilha_slice

import (
	"parte2_benchmark_tests"
	"testing"
)

func BenchmarkPilhaSlicePushPop(b *testing.B) {
	parte2_benchmark_tests.BenchmarkPilhaPushPop(b, &PilhaSlice{}, "../../1984.txt")
}
