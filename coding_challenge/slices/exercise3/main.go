package main

import "fmt"

func main() {

	nums := []float64{10.2, 13.7, 28.9}

	nums = append(nums, 10.1)

	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums)

	n := []float64{88.9, 12.3}

	nums = append(nums, n...)

	fmt.Println(nums)
}
