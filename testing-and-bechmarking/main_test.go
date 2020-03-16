package main

import "testing"

func TestMySum(t *testing.T) {
	x := mySum(2,3)
	if x != 5 {
		t.Errorf("Sum 2 and 3 expected to be 5 but got:%v", x)
	}
}