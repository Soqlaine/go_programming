package main

import "fmt"

func main() {
	var age int = 30
	fmt.Println("Age:", age)

	var name = "Dan"
	//fmt.Println("Your name is:", name)
	_ = name

	var s1 string
	s1 = "Good bye"
	fmt.Println(s1)

	var k int = 6 // not necessary to say the type (int).It is inferred from the literal on the right side of =
	var i = 5
	var j = 5.6
	fmt.Println("i:", i, "j:", j, "k:", k)

	var ii, jj int
	ii = 22
	jj = 33

	ii, jj = jj, ii
	fmt.Println(ii, jj)

	// var car string
	// var cost int

	car, cost := "Audi", 500000
	fmt.Println(car, cost)

	sum := 2 + 5
	fmt.Println(sum)

	var (
		Firstname string
		ages      int
		gender    bool
	)
	_, _, _ = Firstname, ages, gender
}
