package animals_test

import (
	"fmt"
	"sort"

	"github.com/taigrr/animals"
)

func ExampleNames() {
	names := animals.Names()
	fmt.Println(names[0])
	fmt.Println(names[len(names)-1])
	fmt.Println(sort.StringsAreSorted(names))
	// Output:
	// anteater
	// zebra
	// true
}

func ExampleContains() {
	fmt.Println(animals.Contains("cat"))
	fmt.Println(animals.Contains("unicorn"))
	// Output:
	// true
	// false
}

func ExampleStartingWith() {
	fmt.Println(animals.StartingWith("cat"))
	// Output:
	// [cat caterpillar catfish]
}

func ExampleRandomN() {
	sample := animals.RandomN(3)
	fmt.Println(len(sample))
	fmt.Println(animals.Contains(sample[0]))
	// Output:
	// 3
	// true
}

func ExampleAll() {
	count := 0
	for range animals.All() {
		count++
	}
	fmt.Println(count == animals.Count())
	// Output:
	// true
}
