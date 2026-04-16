// Package animals provides a sorted list of common animal names plus a
// handful of convenience helpers for querying and sampling it.
package animals

import (
	"iter"
	"math/rand/v2"
	"slices"
	"strings"
)

var animals = []string{
	"anteater",
	"antelope",
	"aphid",
	"armadillo",
	"asp",
	"ass",
	"baboon",
	"badger",
	"bald eagle",
	"barracuda",
	"bass",
	"basset hound",
	"bat",
	"bearded dragon",
	"beaver",
	"bedbug",
	"bee",
	"bird",
	"bison",
	"black panther",
	"black widow spider",
	"blue jay",
	"blue whale",
	"bobcat",
	"buffalo",
	"butterfly",
	"buzzard",
	"camel",
	"canada lynx",
	"carp",
	"cat",
	"caterpillar",
	"catfish",
	"cheetah",
	"chicken",
	"chimpanzee",
	"chipmunk",
	"cobra",
	"cod",
	"condor",
	"cougar",
	"cow",
	"coyote",
	"crab",
	"crane fly",
	"cricket",
	"crocodile",
	"crow",
	"cuckoo",
	"deer",
	"dinosaur",
	"dog",
	"dolphin",
	"donkey",
	"dove",
	"dragonfly",
	"duck",
	"eagle",
	"eel",
	"elephant",
	"emu",
	"falcon",
	"ferret",
	"finch",
	"fish",
	"flamingo",
	"flea",
	"fly",
	"fox",
	"frog",
	"goat",
	"goose",
	"gopher",
	"gorilla",
	"guinea pig",
	"hamster",
	"hare",
	"hawk",
	"hippopotamus",
	"horse",
	"hummingbird",
	"humpback whale",
	"husky",
	"iguana",
	"impala",
	"kangaroo",
	"lemur",
	"leopard",
	"lion",
	"lizard",
	"llama",
	"lobster",
	"monitor lizard",
	"monkey",
	"moose",
	"mosquito",
	"moth",
	"mouse",
	"mule",
	"octopus",
	"orca",
	"ostrich",
	"otter",
	"owl",
	"ox",
	"oyster",
	"panda",
	"parrot",
	"peacock",
	"pelican",
	"penguin",
	"pheasant",
	"pig",
	"pigeon",
	"polar bear",
	"porcupine",
	"quagga",
	"rabbit",
	"raccoon",
	"rat",
	"rattlesnake",
	"red wolf",
	"rooster",
	"seal",
	"sheep",
	"skunk",
	"sloth",
	"snail",
	"snake",
	"spider",
	"tiger",
	"whale",
	"wolf",
	"wombat",
	"zebra",
}

func init() {
	slices.Sort(animals)
	animals = slices.Compact(animals)
}

// Names returns a copy of the sorted animal name list.
func Names() []string {
	return slices.Clone(animals)
}

// Count returns the number of animals in the list.
func Count() int {
	return len(animals)
}

// Contains reports whether the given name exists in the animal list.
// The comparison is case-sensitive and exact.
func Contains(name string) bool {
	_, ok := slices.BinarySearch(animals, name)
	return ok
}

// All returns an iterator over every animal name in sorted order.
// It is safe to break out of the range early.
func All() iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, name := range animals {
			if !yield(name) {
				return
			}
		}
	}
}

// StartingWith returns a sorted copy of every animal whose name begins with
// the given prefix. An empty prefix returns a full copy equivalent to Names.
func StartingWith(prefix string) []string {
	if prefix == "" {
		return Names()
	}
	start, _ := slices.BinarySearch(animals, prefix)
	var out []string
	for i := start; i < len(animals); i++ {
		if !strings.HasPrefix(animals[i], prefix) {
			break
		}
		out = append(out, animals[i])
	}
	return out
}

// Random returns a random animal name.
func Random() string {
	return animals[rand.IntN(len(animals))]
}

// RandomN returns n distinct random animal names. If n is <= 0, it returns
// nil. If n is greater than Count, it returns all animals in a random order.
func RandomN(n int) []string {
	if n <= 0 {
		return nil
	}
	if n > len(animals) {
		n = len(animals)
	}
	out := slices.Clone(animals)
	rand.Shuffle(len(out), func(i, j int) { out[i], out[j] = out[j], out[i] })
	return out[:n]
}
