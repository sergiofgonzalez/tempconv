# tempconv
> Temperature conversion module between Celsius and Fahrenheit in Go

This is an example of how to create a Go module available to the Go community as described in "The Go Programming Language".

The project is a pure library, and therefore the following project structure is followed:

```
tempconv/
├── .gitignore
├── go.mod
├── LICENSE
├── Makefile
├── README.md
├── tempconv_test.go
└── tempconv.go
```

To use it in a program simply do:

```bash
go get github.com/sergiofgonzalez/tempconv
```

And then, import it and use it in your code:

```go
package main

import (
	"fmt"

	"github.com/sergiofgonzalez/tempconv"
)

func main() {
	fmt.Println("Current temp in Fahrenheit: ", tempconv.CtoF(17))
}
```