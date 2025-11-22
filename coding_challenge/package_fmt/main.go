package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("x type is %T , y type is %T , z type is %T , score type is %T\n", x, y, z, score)
	fmt.Printf("z is %q\n", z)
	fmt.Printf("x value is %v, y value is %v, z value is %v\n", x, y, z)
	fmt.Printf("The type of y is %T and score is %T\n", y, score)

	const n float64 = 1.422349587101
	fmt.Printf("x is %.4f\n", n)
}
