package main

import "fmt"

func myFunc(a int) {

	fmt.Println(a)
}
func main() {

	myFuncVar := myFunc
	fmt.Printf("Type of myFuncVar:%T\n", myFuncVar)
	myFuncVar(8)
}
