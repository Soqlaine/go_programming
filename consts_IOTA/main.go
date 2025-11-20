package main

import (
	"fmt"
)

func main() {
	const days int = 7
	const pi = 3.14

	// there are only boolean constants, integer constants, rune constants

	//multiple grouped const

	const (
		a = 5
		b = 0.1
	)

	const n, m = 4, 5

	const (
		min  = -500
		max1 // get it type and value from the previous constant.
		max2 //get it type and value from the previous constant.
	)

	//CONSTANT RULES

	// 1. you cannot change a constant

	const tmp int = 100
	//temp = 50  // compile time error

	// 2. constant belongs to runtime not compile time

	//const power = math.Pow(2,3) // error function calls belong to runtime

	// 3. you cannot initialize a variable to constant

	t := 5 // variables belongs to runtime

	// const tc = t

	// 4. You can use a function like len() to initialize a const if it has as argument
	// a constant string literal (not a variable)
	// a string literal is an untyped constant

	const l1 = len("Hello") //ok

	str := "Hello"

	// const l2 = len(str) //error str is a var and belongs to runtime

	_, _ = t, str

	// UNTYPED CONSTANTS
	const x = 5
	const y float64 = 1.1

	var v1 int = 5
	var v2 float64 = 1.1

	fmt.Println(x * y)
	// => 5.5, No Error because x is untyped and gets its type when its used first time (float64).

	// fmt.Println(v1 * v2)
	// => Error: invalid operation: v1 * v2 (mismatched types int and float64)
	_, _ = v1, v2

	//IOTA

	// iota is number generator for constants which starts from zero
	// and is incremented by 1 automatically.

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	const (
		c4 = iota * 2
		c5
		c6
	)
	fmt.Println(c4, c5, c6)

	const (
		north = iota
		east
		south
		west
	)
	fmt.Println(north, east, south, west)
}
