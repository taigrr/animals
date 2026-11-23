package animals_test

import (
	"slices"
	"testing"

	"github.com/taigrr/animals"
)

// FuzzContains checks exact membership independently of the binary search used
// by Contains, including names that sort between or outside catalog entries.
func FuzzContains(f *testing.F) {
	names := animals.Names()
	for _, name := range names {
		f.Add(name)
		f.Add(name + "\x00")
	}
	for _, name := range []string{"", "CAT", " cat", "cat ", "catfish\n", "\x00", "\xff", "🐈", "zzzz"} {
		f.Add(name)
	}

	f.Fuzz(func(t *testing.T, name string) {
		want := slices.Contains(names, name)
		if got := animals.Contains(name); got != want {
			t.Fatalf("Contains(%q) = %v, want %v", name, got, want)
		}
	})
}
