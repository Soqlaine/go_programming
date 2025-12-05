package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
}

func main() {

	me := person{name: "Marius", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}

	for index, color := range me.colors {
		fmt.Printf("%d -> %q\n", index, color)
	}

	_ = you
}
