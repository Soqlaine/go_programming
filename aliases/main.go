package main

import "fmt"

func main() {
	var a uint8 = 10
	var b byte //byte is an alias to uint8

	// even though they have different names, byte and uit8 are the same type because they are aliases
	b = a // no error
	_ = b

	// declaring a new alias named second for uint
	// type alias_name = type_name
	type second = uint

	var hour second = 3600
	fmt.Printf("hour type is %T\n", hour)

	fmt.Printf("Minutes in an hour: %d\n", hour/60) // => Minutes in an hour: 60
}
