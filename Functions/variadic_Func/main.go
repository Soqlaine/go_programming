package main

import (
	"fmt"
	"strings"
)

// creating a variadic function
func f1(a ...int) {

	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

// variadic function that modifies one of the arguments passed.

func f2(a ...int) {
	a[0] = 50

	fmt.Printf("%#v\n", a)
}

func sumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.

	for _, v := range a {
		sum += v
		product *= v
	}
	return sum, product

}

func personInformation(age int, names ...string) string {
	fullName := strings.Join(names, " ")
	returnString := fmt.Sprintf("Age: %d, Full Name:%s", age, fullName)
	return returnString

}

func main() {
	f1(2, 3, 4)
	f2(12, 13, 14, 15)

	sum, prod := sumAndProduct(3.4, 7.1)
	fmt.Printf("Sum:%#v, prod:%#v\n", sum, prod)

	fmt.Println(strings.Repeat("#", 20))

	Info := personInformation(24, "Ruvaidha", "Soqlaine", "Aabhar")
	fmt.Println(Info)

}
