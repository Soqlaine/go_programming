package main

import "fmt"

func f1(num uint) (int, int) {
	fact := 1
	sum := 0

	for i := 1; i <= int(num); i++ {
		fact = fact * i

	}

	for n := 0; n <= int(num); n++ {
		sum = sum + n
	}

	return fact, sum
}
func main() {

	factorial, sum := f1(5)
	fmt.Println(factorial, sum)

}
