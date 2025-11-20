package main

import "fmt"

func main() {

	var a int = 10
	var b = 155.5   // type inference deduction
	c := "Gosphers" // short declaration

	_, _, _ = a, b, c

	//a = 3.5 -> error A variable cannot change its type
	//a = b   -> error bcz not allowed to assign a type to another type

	//**Zero values**//

	var value int     // initialized with 0
	var price float64 //initialized with 0.0
	var name string   //initialized with empty string
	var done bool     //initialized with false

	fmt.Println(value, price, name, done)

}
