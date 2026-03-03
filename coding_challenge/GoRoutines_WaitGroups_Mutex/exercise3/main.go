package main

import (
	"fmt"
	"math"
	"sync"
)

// exercise 3 and 4 is almost same

func main() {
	var wg sync.WaitGroup
	wg.Add(50)

	for i := 100.; i <= 150.; i++ {
		go func(n float64, wg *sync.WaitGroup) {
			x := math.Sqrt(n)
			fmt.Printf("%.2f\n", x)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()

	// var wg sync.WaitGroup

	// wg.Add(50)

	// // IMPORTANT: i starts from 100. not 100
	// // i is float64, not int. math.Sqrt() takes in a float64.
	// for i := 100.; i < 150.; i++ {
	// 	go func(f float64, wg *sync.WaitGroup) {
	// 		x := math.Sqrt(f)
	// 		fmt.Printf("Square root of %.2f is %.6f\n", f, x)
	// 		wg.Done()
	// 	}(i, &wg)
	// }

	// wg.Wait()

}
