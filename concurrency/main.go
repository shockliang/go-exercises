package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("OS\t\t",runtime.GOOS)
	fmt.Println("ARCH\t\t",runtime.GOARCH)
	fmt.Println("CPUs\t\t",runtime.NumCPU())
	fmt.Println("Goroutines\t",runtime.NumGoroutine())

	go foo()
	bar()

	fmt.Println("CPUs\t\t",runtime.NumCPU())
	fmt.Println("Goroutines\t",runtime.NumGoroutine())
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
