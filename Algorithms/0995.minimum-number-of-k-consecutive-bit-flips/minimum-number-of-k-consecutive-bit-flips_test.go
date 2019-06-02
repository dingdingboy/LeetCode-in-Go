package problem0995

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	A   []int
	K   int
	ans int
}{

	{
		[]int{0, 0, 0, 1, 0, 1, 1, 0},
		3,
		3,
	},

	{
		[]int{0, 1, 0},
		1,
		2,
	},

	{
		[]int{1, 1, 0},
		2,
		-1,
	},

	// 可以有多个 testcase
}

func Test_minKBitFlips(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, minKBitFlips(tc.A, tc.K), "输入:%v", tc)
	}
}

func Benchmark_minKBitFlips(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minKBitFlips(tc.A, tc.K)
		}
	}
}
