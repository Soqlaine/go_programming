package main

import "fmt"

func sum(a ...int) (sums int) {

	for _, v := range a {
		sums += v
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3))
}
