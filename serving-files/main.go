package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	io.WriteString(w, `
		<!-- not serving from our server -->
		<img src="https://i.insider.com/5df773cefd9db21a1c58b0c4?width=2500&format=jpeg&auto=webp" >`)
}
