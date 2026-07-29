package animals_test

import (
	"fmt"

	"github.com/taigrr/animals"
)

func ExampleNames() {
	names := animals.Names()
	fmt.Println(names[:3])
	// Output:
	// [anteater antelope aphid]
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

func ExampleAll() {
	var firstThree []string
	for name := range animals.All() {
		firstThree = append(firstThree, name)
		if len(firstThree) == 3 {
			break
		}
	}
	fmt.Println(firstThree)
	// Output:
	// [anteater antelope aphid]
}
