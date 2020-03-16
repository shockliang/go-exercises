package main

func main() {

}

func mySum(xi ...int) int {
	sum := 0
	for _, y:= range xi {
		sum += y
	}

	return sum
}