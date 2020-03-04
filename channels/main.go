package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make (<-chan int)	// receive
	cs := make(chan <- int) // send

	fmt.Printf("%T\n",c)
	fmt.Printf("%T\n",cr)
	fmt.Printf("%T\n",cs)

	// general to specific assigns
	cr = c
	cs = c

	fmt.Printf("cr\t%T\n",cr)
	fmt.Printf("cs\t%T\n",cs)

	// general to specific converts
	fmt.Printf("c\t%T\n", (<-chan int)(c))
	fmt.Printf("c\t%T\n", (chan<- int)(c))
}