package pilha_list

import (
	"parte2_benchmark_tests"
	"testing"
)

func BenchmarkPilhaListPushPop(b *testing.B) {
	parte2_benchmark_tests.BenchmarkPilhaPushPop(b, &PilhaList{}, "../../1984.txt")
}
