package animals

import (
	"strconv"
	"testing"
)

func BenchmarkRandomN(b *testing.B) {
	for _, count := range []int{1, 5, Count()} {
		b.Run(strconv.Itoa(count), func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				RandomN(count)
			}
		})
	}
}
