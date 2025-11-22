package main

import (
	"fmt"
	"strconv"
)

func main() {

	i, err := strconv.Atoi("55")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if x, err := strconv.Atoi("45"); err == nil {
		fmt.Println(x)
	} else {
		fmt.Println(err)
	}
}
