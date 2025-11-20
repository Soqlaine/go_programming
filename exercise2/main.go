/*package main

import "fmt"

func main() {
	var1 := 20
	var2 := 40
	fmt.Println("Variable 1 value is", var1)
	fmt.Println("Variable 2 value is", var2)
}*/

package main

import "fmt"

func main() {
	SayHello()
}
func SayHello() {
	fmt.Println("Hello package v1.0.0!")

	var x int = 10
	_ = x
}
