package main

import "fmt"

func main() {

	type grades struct {
		grade  int
		course string
	}

	type Person struct {
		name       string
		age        int
		gradesInfo grades
	}

	me := Person{
		"Ruvaidha",
		25,
		grades{19, "Python"},
	}

	you := Person{
		"Soqlaine",
		15,
		grades{27, "Java"},
	}

	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)

	you.gradesInfo.grade = 98
	you.gradesInfo.course = "Golang"

	fmt.Printf("%v\n", me)
	fmt.Printf("%v\n", you)

}
