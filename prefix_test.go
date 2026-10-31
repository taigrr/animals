package animals

import (
	"slices"
	"strings"
	"testing"
)

// Check the binary-search implementation against an independent linear scan.
func checkStartingWith(t *testing.T, prefix string) {
	t.Helper()
	var want []string
	for _, name := range Names() {
		if strings.HasPrefix(name, prefix) {
			want = append(want, name)
		}
	}
	if got := StartingWith(prefix); !slices.Equal(got, want) {
		t.Fatalf("StartingWith(%q) = %q, want %q", prefix, got, want)
	}
}

func TestStartingWithEveryPrefix(t *testing.T) {
	prefixes := map[string]struct{}{
		"": {}, "\x00": {}, "\xff": {}, "CAT": {}, "é": {},
		"cata": {}, "cat\x00": {}, "zebras": {},
	}
	for _, name := range Names() {
		for end := 1; end <= len(name); end++ {
			prefixes[name[:end]] = struct{}{}
		}
		// Exercise insertion points immediately after complete names.
		prefixes[name+"\x00"] = struct{}{}
	}
	for prefix := range prefixes {
		t.Run(prefix, func(t *testing.T) {
			checkStartingWith(t, prefix)
		})
	}
}

func FuzzStartingWith(f *testing.F) {
	for _, prefix := range []string{"", "cat", "cata", "b", "zebra", "zzz", "CAT", "\x00", "\xff", "é"} {
		f.Add(prefix)
	}
	f.Fuzz(func(t *testing.T, prefix string) {
		checkStartingWith(t, prefix)
	})
}
