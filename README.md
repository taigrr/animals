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

	// Check if an animal exists
	fmt.Println(animals.Contains("cat")) // true

	// Get the count
	fmt.Println(animals.Count()) // 130
}
```

## License

0BSD
