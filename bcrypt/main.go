package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password2234`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bs)

	loginPw := `password2234`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPw))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("you are logged in")
}
