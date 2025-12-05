package main

import "fmt"

func sum(a ...int) int {
	sums := 0
	for i := 0; i < len(a); i++ {
		sums += a[i]
	}
	return sums
}

func sums(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}
func main() {
	sumOver := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(sumOver)
}
