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

func ExampleMySum() {
	fmt.Println(MySum(2, 3))
	// Output:
	// 5
}

func TestMySum(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 40}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := MySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.data, " Got", x)
		}
	}

}
