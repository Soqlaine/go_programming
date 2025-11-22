package main

import "fmt"

func main() {

	// exercise-1

	// for i := 1; i <= 50; i++ {
	// 	if i%7 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//exercise-2

	// for i := 1; i <= 50; i++ {
	// 	if i%7 != 0 {
	// 		continue
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	//exercise-3

	// count := 0
	// for i := 1; i <= 50; i++ {
	// 	if i%7 == 0 {
	// 		fmt.Println(i)
	// 		count++
	// 	}

	// 	if count == 3 {
	// 		break
	// 	}
	// }

	//exercise-4

	// for i := 1; i <= 500; i++ {
	// 	if i%5 == 0 && i%7 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//exercise-5

	birthyear := 2000
	currentyear := 2025

	for i := birthyear; i <= currentyear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

}
