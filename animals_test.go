package animals

import (
	"sort"
	"testing"
)

func TestNames(t *testing.T) {
	names := Names()
	if len(names) == 0 {
		t.Fatal("expected non-empty animal list")
	}
	if !sort.StringsAreSorted(names) {
		t.Fatal("expected sorted list")
	}
}

func TestNamesReturnsCopy(t *testing.T) {
	names := Names()
	original := names[0]
	names[0] = "zzz-mutated"
	fresh := Names()
	if fresh[0] != original {
		t.Fatal("Names() should return a copy; mutation leaked to original")
	}
}

func TestNamesContains(t *testing.T) {
	names := Names()
	found := false
	for _, name := range names {
		if name == "cat" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("expected 'cat' in list")
	}
}

func TestCount(t *testing.T) {
	count := Count()
	names := Names()
	if count != len(names) {
		t.Fatalf("Count() = %d, len(Names()) = %d", count, len(names))
	}
	if count == 0 {
		t.Fatal("expected non-zero count")
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"cat", true},
		{"dog", true},
		{"zebra", true},
		{"anteater", true},
		{"unicorn", false},
		{"", false},
		{"CAT", false},
	}
	for _, tt := range tests {
		if got := Contains(tt.name); got != tt.want {
			t.Errorf("Contains(%q) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestRandom(t *testing.T) {
	got := Random()
	if got == "" {
		t.Fatal("Random() returned empty string")
	}
	if !Contains(got) {
		t.Fatalf("Random() returned %q which is not in the list", got)
	}
}

func TestRandomDistribution(t *testing.T) {
	seen := make(map[string]bool)
	for range 100 {
		seen[Random()] = true
	}
	if len(seen) < 5 {
		t.Fatalf("Random() returned only %d unique animals in 100 calls", len(seen))
	}
}
