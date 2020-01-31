package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{
					Number: "CSCI-40",
					Name:   "Introduction to Programming in Go",
					Units:  "4",
				},
				{
					Number: "CSCI-130",
					Name:   "Introduction to Web Programming with Go",
					Units:  "4",
				},
				{
					Number: "CSCI-140",
					Name:   "Mobile Apps Using Go",
					Units:  "4",
				},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{
					Number: "CSCI-50",
					Name:   "Advanced Go",
					Units:  "5",
				},
				{
					Number: "CSCI-190",
					Name:   "Advanced Web Programming with Go",
					Units:  "5",
				},
				{
					Number: "CSCI-191",
					Name:   "Advanced Mobile Apps with Go",
					Units:  "5",
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)

	if err != nil {
		log.Fatalln(err)
	}
}
