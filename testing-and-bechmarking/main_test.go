package main

import (
	"fmt"
	"testing"
)

func ExampleMySum() {
	fmt.Println(MySum(2,3))
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
