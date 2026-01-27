package main

import "fmt"

// func swap(a float64, b float64) (aa float64, bb float64) {
// 	c := a
// 	a = b
// 	b = c

// 	return a, b

// }

func swap_Pointer(a, b *float64) {
	*a, *b = *b, *a
}

func main() {

	x, y := 5.5, 8.8
	// xx, yy := swap(x, y)
	// fmt.Println(xx, yy)

	swap_Pointer(&x, &y)
	fmt.Printf("x is %v , y is %v", x, y)

}
