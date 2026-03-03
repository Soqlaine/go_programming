package main

import (
	"fmt"
	"sync"
)

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go sayHello("Mr. Wick", &wg)

	wg.Wait()
}

// Execution flow (step-by-step)

// main() starts

// wg created (counter = 0)

// wg.Add(1) → counter = 1

// go sayHello(...) → goroutine starts

// main() hits wg.Wait() → waits

// sayHello() prints message

// defer wg.Done() executes → counter = 0

// wg.Wait() unblocks

// main() exits

// Program ends safely
