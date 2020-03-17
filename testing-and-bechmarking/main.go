// My main package
package main

func main() {

}

// Sum adds an unlimited number of values of type int
func MySum(xi ...int) int {
	sum := 0
	for _, y:= range xi {
		sum += y
	}

	return sum
}