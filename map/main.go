package main

import "fmt"

func main() {
	//colors := map[string]string {
	//	"red": "#ff0000",
	//	"olive": "#808000",
	//}

	//var colors map[string]string

	colors := make(map[string]string)
	colors["red"] = "#ff0000"
	colors["white"] = "#ffffff"

	delete(colors, "red")

	fmt.Println(colors)
}
