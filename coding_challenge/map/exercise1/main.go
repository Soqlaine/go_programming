package main

import "fmt"

func main() {

	// var m1 map[string]string

	// fmt.Printf("Type of M1 is: %s\n", m1)

	// m2 := map[int]string{10: "Abba"}

	// fmt.Println(m2[10])
	// fmt.Println(m2[9])

	// 1.
	var m1 map[float64]bool
	fmt.Printf("m1 type: %T, m1 value: %#v\n", m1, m1)

	// 2.
	m2 := map[int]string{1: "Sting", 2: "Queen"}

	// 3.
	m2[10] = "Abba"

	// 4
	fmt.Println(m2[2])   // existing key
	fmt.Println(m2[100]) // non-existing key

}
