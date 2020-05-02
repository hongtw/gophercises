package arraypacking

import (
	"math/rand"
	"testing"
)

type sample struct {
	v   []int
	ans int
}

func generateRuningSamples(n int) [][]int {
	samples := make([][]int, n)
	for i := 0; i < n; i++ {
		v := make([]int, 4)
		for j := 0; j < 4; j++ {
			v[j] = rand.Intn(256)
		}
		samples[i] = v
	}
	return samples
}

var benchmarkSamples = generateRuningSamples(10000)

func BenchmarkV1(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		for _, value := range benchmarkSamples {
			arrayPacking(value)
		}
	}
}

func BenchmarkV2(gb *testing.B) {
	for i := 0; i < gb.N; i++ {
		for _, value := range benchmarkSamples {
			arrayPackingV2(value)
		}
	}
}

func Test_arrayPacking(t *testing.T) {
	tests := []sample{
		{[]int{24, 85, 0}, 21784},
		{[]int{187, 99, 42, 43}, 724198331},
	}
	for i, test := range tests {
		if got := arrayPacking(test.v); got != test.ans {
			t.Errorf("[test-%v] v2: got %v, need %v \n", i, got, test.ans)
		}
		if got := arrayPackingV2(test.v); got != test.ans {
			t.Errorf("[test-%v] v2: got %v, need %v \n", i, got, test.ans)
		}
	}
}
