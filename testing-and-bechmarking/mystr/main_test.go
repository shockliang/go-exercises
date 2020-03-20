package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s =Cat(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "expected", "Shaken not stirred")
	}
}

func TestJoin(t *testing.T) {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	s = Join(xs)
	if s != "Shaken not stirred" {
		t.Error("got", s, "expected", "Shaken not stirred")
	}
}

func ExampleCat() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Cat(xs))
	// Output
	// Shaken not stirred
}

func ExampleJoin() {
	s := "Shaken not stirred"
	xs := strings.Split(s, " ")
	fmt.Println(Join(xs))
	// Output
	// Shaken not stirred
}

const s = "Contrary to popular belief, " +
	"Lorem Ipsum is not simply random text. " +
	"It has roots in a piece of classical Latin " +
	"literature from 45 BC, making it over 2000 years old." +
	" Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia" +
	", looked up one of the more obscure Latin words, consectetur, " +
	"from a Lorem Ipsum passage, and going through the cites of the word in classical literature," +
	" discovered the undoubtable source. " +
	"Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of de Finibus Bonorum et Malorum " +
	"(The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, " +
	"very popular during the Renaissance. " +
	"The first line of Lorem Ipsum, Lorem ipsum dolor sit amet.., comes from a line in section 1.10.32."

var xs = strings.Split(s, " ")

func BenchmarkCat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}
}

func BenchmarkJoin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Join(xs)
	}
}