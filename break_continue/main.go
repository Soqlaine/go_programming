package main

import "fmt"

func main() {

	//** CONTINUE STATEMENT **//

	// It works just the same as in C,  Java or Python.
	// The continue statement rejects all the remaining statements in the current iteration of the loop
	// and moves the control back to the top of the loop.

	// printing even numbers less than or equal to 10
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}

	// **BREAK STATEMENT **//

	// It is used to terminate the innermost for or switch statement.
	// It works just the same as in C,  Java or Python.

	// finding 10 numbers divisible by 13

	fmt.Println("**********************")
	count := 0

	for x := 1; true; x++ {
		if x%13 != 0 {
			continue
		} else {
			count++
			fmt.Printf("%d is divisible by 13\n", x)
		}
		if count == 10 {
			break
		}
	}

	// for i := 0; true; i++ {
	// 	if i%13 == 0 {
	// 		fmt.Printf("%d is divisible by 13\n", i)
	// 		count++
	// 	}
	// 	if count == 10 {
	// 		break
	// 	}
	// }
	fmt.Println("Just a message after a loop")

}
