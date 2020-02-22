package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		age:   32,
	}
	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		age:   27,
	}

	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}
