package main

import "fmt"

func main() {

	type duration int

	var hour duration = 3600

	fmt.Printf("The value of hour is %v\n", hour)

	fmt.Printf("The type of hour is %T\n", hour)

}
