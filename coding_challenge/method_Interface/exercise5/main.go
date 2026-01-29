package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenseno string
	brand     string
}

func (c car) License() string {
	return c.licenseno
}

func (c car) Name() string {
	return c.brand
}

func main() {

	var v vehicle = car{"ANVM00231", "Inova"}
	fmt.Println(v.License())
	fmt.Println(v.Name())

}
