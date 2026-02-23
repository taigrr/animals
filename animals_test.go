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

func TestNamesContains(t *testing.T) {
	names := Names()
	found := false
	for _, n := range names {
		if n == "cat" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("expected 'cat' in list")
	}
}
