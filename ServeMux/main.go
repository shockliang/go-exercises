package main

import (
	"io"
	"net/http"
)

type hotdog int

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}


func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty kitty kitty")
}

func main() {
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
