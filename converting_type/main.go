package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := 5
	b := 3.2

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	c := float64(a)
	d := int32(b)

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	// a = int(float64(a)* b)
	// fmt.Println(a)

	//In Go types with different names are different types.
	var a1 int = 5   // same size as int64 or int32 (platform specific)
	var b1 int64 = 2 // int and int64 are not the same type

	// a1 = b1 // error: cannot use b (type int64) as type int in assignment

	a1 = int(b1) // converting int64 to int (explicit conversion required)

	// preventing unused variable error
	_ = a1

	//** CONVERTING NUMBERS TO STRINGS AND STRINGS TO NUMBERS **//

	s := string(99) // int to rune (Unicode code point)
	fmt.Println(s)  // => 99, the ascii code for symbol c

	fmt.Println(string(34234)) // => 34234 is the unicode code point for 薺

	// we cannot convert a float to a string similar to an int to a string
	// s1 := string(65.1) // error

	// converting float to string
	var myStr = fmt.Sprintf("%f", 5.12)
	fmt.Println(myStr)

	// converting int to string
	var myStr1 = fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1) // => 34234

	// converting string to float
	var result, err = strconv.ParseFloat("3.14", 64)
	if err == nil {
		fmt.Println("Type: %T, Value: %v\n", result, result)
	} else {
		fmt.Println("Cannot convert into float64")
	}

	// Atoi(string to int) and Itoa(int to string).

	i, err := strconv.Atoi("50")
	s = strconv.Itoa(20)

	fmt.Printf("i Type is %T and i value is %v\n", i, i)

	fmt.Printf("s Type is %T and s value is %v\n", s, s)

}
