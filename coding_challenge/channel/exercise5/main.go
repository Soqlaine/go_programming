package main

import "fmt"

func main() {
	ch := make(chan int)
	for i := 1; i <= 50; i++ {
		go func(x int) {
			ch <- x * x
		}(i)
		fmt.Println(<-ch)
	}

}
