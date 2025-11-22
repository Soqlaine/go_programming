package main

import "fmt"

func main() {

	const distance = 149600000000
	const speed = 299792458

	time := distance / speed
	fmt.Println("The time taken for the sunlight to reach the earth is ", time, "seconds")
}
