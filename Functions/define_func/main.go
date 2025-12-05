package main

import (
	"fmt"
	"math"
)

func f1() {

	fmt.Println("This is f1() function")
}

func f2(a int, b int) {
	fmt.Println("Sum: ", a+b)
}

func f3(a, b, c int, d, e float64, s string) {

	fmt.Println(a, b, c, d, e, s)
}

func f4(a float64) float64 {
	return math.Pow(a, a)
}

func f5(a, b int) (int, int) {
	return a + b, a * b
}

func sum(a, b int) int {
	// fmt.Println("s:", s)

	return a + b
}

func main() {

	f1()
	f2(2, 3)
	f3(3, 4, 5, 2., 5.5, "ss")
	fmt.Println(f4(2.2))
	a, b := f5(3, 7)
	fmt.Printf("a:%d ,b:%d\n", a, b)
	// fmt.Println(f5(3, 7))

	ss := sum(4, 5)
	fmt.Println(ss)

}
