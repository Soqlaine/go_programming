package main

import "fmt"

func main() {

	// if condition_that_evaluates_to_boolean{
	//      perform action1
	// }else if condition_that_evaluates_to_boolean{
	//      perform action2
	// }else{
	//      perform action3
	// }

	price, inStock := 100, true

	if price >= 80 {
		fmt.Println("Too expensive")
	}

	if price <= 100 && inStock == true {
		fmt.Println("Buy it")
	}

	// In Go there is not such a thing like the Truthiness of a variable.
	// Error:
	// if price {
	// 	fmt.Println("We have price!")
	// }

	if price < 100 {
		fmt.Println("Its cheap")
	} else if price == 100 {
		fmt.Println("On the edge")
	} else {
		fmt.Println("Its expensive")
	}
}
