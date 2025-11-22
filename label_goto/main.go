package main

import "fmt"

func main() {

	//** LABEL STATEMENT **//

	// declaring a variable
	// there is no conflict name between variable and label

	outer := 19
	_ = outer

	people := [5]string{"Helen", "Mark", "Brenda", "Antonia", "Michael"}
	friends := [2]string{"Mark", "Antonia"}

outer: //label, it doesn't conflict with other names
	// iterating over the array.

	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend: %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")

	// **GOTO STATEMENT **//

	//the following piece of code creates a loop like a for statement does

	i := 0

loop:

	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//  goto todo //ERROR it's not permitted to jump over the declaration of x
	//  x := 5
	// todo:
	//  fmt.Println("something here")

}
