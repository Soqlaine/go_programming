package main

import (
	"fmt"
	"strconv"
)

// func main() {

// 	// var i int = 3
// 	// var f float64 = 3.2

// 	// var i1 = float64(i)
// 	// var f1 = int(f)

// 	// fmt.Printf("Type of i1 is %T and type of f1 is %T\n", i1, f1)

// 	var i = 3
// 	var f = 3.2
// 	var s1, s2 = "3.14", "5"

// 	s := strconv.Itoa(i)
// 	fmt.Println(s)

// 	ss, err := strconv.Atoi(s2)
// 	if err == nil {
// 		fmt.Println(ss)
// 	} else {
// 		fmt.Println("cannot convert")
// 	}

// 	fval := fmt.Sprintf("%f", f)
// 	fmt.Println(fval)

// 	stof, err := strconv.ParseFloat(s1, 64)
// 	if err == nil {
// 		fmt.Println(stof)
// 	} else {
// 		fmt.Println("cant convert")
// 	}
// }

func main() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// 1. int to string
	s := strconv.Itoa(i)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// 2. string to int
	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// 3. float64 to string
	ss1 := fmt.Sprintf("%f", f)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1)

	// 4. string to float64
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}
}
