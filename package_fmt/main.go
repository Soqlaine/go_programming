package main

import "fmt"

func main() {

	//fmt.Println()

	fmt.Println("Hello Go World")

	var name, age = "Andrei", 35

	fmt.Println(name, "is", age, "years old.")

	//fmt.Printf()  prints out to stdout according to a format specifier called verb. it doesn't add a new line(\n).

	//VERBS:
	// %d -> decimal
	// %f -> float
	// %s -> string
	// %q -> double-quoted string
	// %v -> value (any)
	// %#v -> a Go-syntax representation of the value
	// %T -> value Type
	// %t -> bool (true or false)
	// %p -> pointer (address in base 16, with leading 0x)
	// %c -> char (rune) represented by the corresponding Unicode code point

	a, b, c := 10, 15.5, "Gosphers"
	grades := []int{10, 20, 30}
	// _ = grades

	fmt.Printf("a is %d, b is %f, c is %s\n", a, b, c)

	fmt.Printf("%q\n", c)

	fmt.Printf("%v\n", grades)

	fmt.Printf("b is a type of %T  and c is a type of %T and grades is a type of %T\n", a, b, grades)

	fmt.Printf("The address of a: %p\n", &a)

	fmt.Printf("%c and %c\n", 100, 51012)

	const pi = 3.1415159265359
	fmt.Printf("pi is %.4f\n", pi)

	// %b -> base2
	// %x -> base16

	fmt.Printf("255 in base 2 is %b\n", 255)

	fmt.Printf("101 in base 16 is %x\n", 101)

	//fmt.Sprintf() returns a string, Uses the same verbs of fmt.Printf()

	s := fmt.Sprintf("a is %d, b is %f, c is %s\n", a, b, c)
	fmt.Println(s)

}
