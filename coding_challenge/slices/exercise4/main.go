package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]

	sum, prod := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 to 10 numbers!")
	} else {
		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue
			}
			sum += num
			prod *= num
		}
	}

	fmt.Printf("Sum:%v,Prod:%v\n", sum, prod)

}
