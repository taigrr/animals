# animals

A simple Go package providing a sorted list of animal names.

## Install

```bash
go get github.com/taigrr/animals
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/taigrr/animals"
)

func main() {
	// Get all animal names (sorted)
	names := animals.Names()
	fmt.Println(names)

	// Get a random animal
	fmt.Println(animals.Random())

	// Get 5 distinct random animals
	fmt.Println(animals.RandomN(5))

	// Filter by prefix
	fmt.Println(animals.StartingWith("cat")) // [cat caterpillar catfish]

	// Check if an animal exists
	fmt.Println(animals.Contains("cat")) // true

	// Get the count
	fmt.Println(animals.Count())

	// Iterate without allocating a slice (Go 1.23+ range-over-func)
	for name := range animals.All() {
		fmt.Println(name)
	}
}
```

## License

0BSD
