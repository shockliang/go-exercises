// My main package
package main

import (
	"fmt"
	"github.com/shockliang/go-exercises/testing-and-bechmarking/saying"
)

func main() {

	fmt.Println(saying.Greet("Jame"))
}

// Sum adds an unlimited number of values of type int
func MySum(xi ...int) int {
	sum := 0
	for _, y:= range xi {
		sum += y
	}

	return sum
}