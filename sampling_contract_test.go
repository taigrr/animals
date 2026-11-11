package animals

import (
	"slices"
	"strconv"
	"testing"
)

func TestRandomNPreservesSourceAndOwnsResult(t *testing.T) {
	original := Names()
	for _, size := range []int{1, Count() / 2, Count(), Count() + 1} {
		t.Run(strconv.Itoa(size), func(t *testing.T) {
			sample := RandomN(size)
			if !slices.Equal(Names(), original) {
				t.Fatal("sampling changed the source list")
			}
			retained := slices.Clone(sample)
			other := RandomN(Count())
			for index := range other {
				other[index] = "not an animal"
			}
			if !slices.Equal(sample, retained) {
				t.Fatal("separate samples share storage")
			}
			for index := range sample {
				sample[index] = "not an animal"
			}
			if !slices.Equal(Names(), original) {
				t.Fatal("mutating a sample changed the source list")
			}
		})
	}
}

func TestAllReusableAfterEarlyStop(t *testing.T) {
	sequence := All()
	for range sequence {
		break
	}
	for attempt := range 2 {
		if got := slices.Collect(sequence); !slices.Equal(got, Names()) {
			t.Fatalf("iteration %d did not restart with the complete sorted list", attempt)
		}
	}
}

func TestRandomNNonPositiveReturnsNil(t *testing.T) {
	for _, size := range []int{-1, 0} {
		if got := RandomN(size); got != nil {
			t.Errorf("RandomN(%d) = %v, want nil", size, got)
		}
	}
}
