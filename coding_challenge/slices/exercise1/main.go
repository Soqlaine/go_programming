package main

import "fmt"

func main() {

	// exercise - 1

	a := []string{"Aabra", "Ka", "Daabra"}

	for i, val := range a {
		fmt.Println("Index: ", i, "value: ", val)
	}

}
