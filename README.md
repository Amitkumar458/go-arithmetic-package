# arithmetic-go

A simple Go package for basic arithmetic operations.

## Installation

```sh
go get github.com/Amitkumar458/go-arithmetic-package

package main

import (
	"fmt"
	"github.com/Amitkumar458/go-arithmetic-package"
)

func main() {
	fmt.Println(arithmetic.Add(5, 3))      // 8
	fmt.Println(arithmetic.Subtract(10, 4)) // 6
	fmt.Println(arithmetic.Multiply(3, 4))  // 12
	result, err := arithmetic.Divide(8, 2)
	if err == nil {
		fmt.Println(result) // 4
	}
}