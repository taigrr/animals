package animals

import (
	"slices"
	"strings"
	"testing"
)

func TestNames(t *testing.T) {
	names := Names()
	if len(names) == 0 {
		t.Fatal("expected non-empty animal list")
	}
	if !slices.IsSorted(names) {
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
	if !slices.Contains(Names(), "cat") {
		t.Fatal("expected 'cat' in list")
	}
}

func TestNamesNoDuplicates(t *testing.T) {
	names := Names()
	seen := make(map[string]struct{}, len(names))
	for _, n := range names {
		if _, dup := seen[n]; dup {
			t.Fatalf("duplicate animal in list: %q", n)
		}
		seen[n] = struct{}{}
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

func TestAll(t *testing.T) {
	var collected []string
	for name := range All() {
		collected = append(collected, name)
	}
	if !slices.Equal(collected, Names()) {
		t.Fatal("All() should yield every animal in sorted order")
	}
}

func TestAllEarlyBreak(t *testing.T) {
	count := 0
	for range All() {
		count++
		if count == 3 {
			break
		}
	}
	if count != 3 {
		t.Fatalf("expected to stop after 3 items, saw %d", count)
	}
}

func TestStartingWith(t *testing.T) {
	tests := []struct {
		prefix  string
		wantMin int
		check   func([]string) error
	}{
		{"b", 10, func(out []string) error {
			for _, s := range out {
				if !strings.HasPrefix(s, "b") {
					t.Errorf("StartingWith(\"b\") returned %q", s)
				}
			}
			return nil
		}},
		{"cat", 3, func(out []string) error {
			if !slices.Contains(out, "cat") {
				t.Error("StartingWith(\"cat\") missing 'cat'")
			}
			if !slices.Contains(out, "caterpillar") {
				t.Error("StartingWith(\"cat\") missing 'caterpillar'")
			}
			return nil
		}},
		{"zzz", 0, nil},
	}
	for _, tt := range tests {
		got := StartingWith(tt.prefix)
		if len(got) < tt.wantMin {
			t.Errorf("StartingWith(%q) returned %d results, want at least %d", tt.prefix, len(got), tt.wantMin)
		}
		if !slices.IsSorted(got) {
			t.Errorf("StartingWith(%q) result not sorted", tt.prefix)
		}
		if tt.check != nil {
			_ = tt.check(got)
		}
	}
}

func TestStartingWithEmpty(t *testing.T) {
	got := StartingWith("")
	if !slices.Equal(got, Names()) {
		t.Fatal("StartingWith(\"\") should equal Names()")
	}
}

func TestRandomN(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
		{5, 5},
		{Count(), Count()},
		{Count() + 100, Count()},
	}
	for _, tt := range tests {
		got := RandomN(tt.n)
		if len(got) != tt.want {
			t.Errorf("RandomN(%d) returned %d items, want %d", tt.n, len(got), tt.want)
		}
		seen := make(map[string]struct{}, len(got))
		for _, name := range got {
			if _, dup := seen[name]; dup {
				t.Errorf("RandomN(%d) returned duplicate %q", tt.n, name)
			}
			seen[name] = struct{}{}
			if !Contains(name) {
				t.Errorf("RandomN(%d) returned unknown animal %q", tt.n, name)
			}
		}
	}
}
