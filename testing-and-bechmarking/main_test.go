package main

import (
	"fmt"
	"github.com/shockliang/go-exercises/testing-and-bechmarking/saying"
	"testing"
)

func TestGreet(t *testing.T) {
	s := saying.Greet("James")
	if s != "Hello and welcome James" {
		t.Error("got", s, "expected", "Hello and welcome James")
	}
}

func ExampleGreet() {
	fmt.Println(saying.Greet("James"))
	// Output:
	// Hello and welcome James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		saying.Greet("James")
	}
}
