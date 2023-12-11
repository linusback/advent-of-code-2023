package day11

import "testing"

func Benchmark_Solve(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Solve()
		if err != nil {
			b.Error(err)
		}
	}
}
