package main

import "fmt"

func main() {
	var c1 chan float64 //birectional channel

	c2 := make(<-chan rune) // receive only channel
	c3 := make(chan<- rune) //send only channel

	c4 := make(chan int, 10) //bidirectional channel with capacity

	fmt.Printf("%T\n, %T\n, %T\n, %T\n,", c1, c2, c3, c4)

	//example
	c4 <- 12
	n := <-c4

	fmt.Println(n)

}
