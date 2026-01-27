package main

import "fmt"

func main() {

	x := 10.10
	fmt.Println(&x)

	ptr := &x
	fmt.Printf("Type of ptr %T and value of x %v\n", ptr, ptr)

	fmt.Printf("Address of the pointer:%p\n ", &ptr)

	fmt.Printf("Value of x through the pointer:%f\n", *ptr)
}
