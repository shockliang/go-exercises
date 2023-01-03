package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandInt(min, max int) <-chan int {
	out := make(chan int, 3)

	go func() {
		for {
			out <- rand.Intn(max-min+1) + min
		}
	}()

	return out
}

func generateRandIntNum(count, min, max int) <-chan int {
	out := make(chan int, 1)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Intn(max-min+1) + min
		}
		close(out)
	}()

	return out
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randInt := generateRandInt(1, 100)

	fmt.Println("generate rand int infinite")

	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)
	fmt.Println(<-randInt)

	fmt.Println("generate rand int by count")
	randIntRange := generateRandIntNum(2, 1, 10)
	for i := range randIntRange {
		fmt.Println(i)
	}

	fmt.Println("out of channel range")
	fmt.Println(<-randIntRange)
	fmt.Println(<-randIntRange)
	fmt.Println(<-randIntRange)

	fmt.Println("generate rand int by count using infinite loop to print out")
	randIntNum := generateRandIntNum(3, 1, 10)
	for {
		n, open := <-randIntNum
		if !open {
			break
		}

		fmt.Println(n)
	}

}
