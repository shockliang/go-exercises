package main

import "fmt"

func main() {
	name := "Shock"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>HELLO WORLD</title>
	</head>
	<body>
		<h1>` + name + `</h1
	</body>
	</head>
	</html>
	`
	fmt.Println(tpl)
}
