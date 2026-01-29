package main

import "fmt"

type empty interface {
}

func main() {

	var v empty
	fmt.Printf("type: %T\n", v)

	v = 5
	fmt.Printf("Type: %T\n", v)

	v = 5.9
	fmt.Printf("Type: %T\n", v)

	v = []int{1, 2}
	fmt.Printf("Type: %T\n", v)

	v = append(v.([]int), 10)
	fmt.Printf("values: %v\n", v)

}
