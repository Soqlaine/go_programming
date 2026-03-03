package main

import (
	"fmt"
	"sync"
)

func sum(s1, s2 float64, wg *sync.WaitGroup) {
	sums := s1 + s2
	fmt.Printf("%.2f\n", sums)
	wg.Done()
}
func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go sum(2.3, 7.9, &wg)
	go sum(5, 7, &wg)
	go sum(12.334, 1.456, &wg)

	wg.Wait()
}
