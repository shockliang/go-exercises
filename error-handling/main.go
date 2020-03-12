package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		// err happened open no-file.txt: no such file or directory
		//fmt.Println("err happened", err)
		// 2020/03/11 22:14:00 err happened open no-file.txt: no such file or directory
		//log.Println("err happened", err)
		//log.Fatalln(err)
		//log.Panic(err)
		panic(err)
	}

	defer f2.Close()
	fmt.Println("check the log.txt file in the directory")
}
