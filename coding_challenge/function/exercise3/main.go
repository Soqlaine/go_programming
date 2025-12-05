package main

import (
	"fmt"
	"os"
	"strconv"
)

func myFunc(a string) int {

	s, err := strconv.Atoi(a)

	if err != nil {
		fmt.Printf("Cannot conevrt %q to int", a)
		os.Exit(1)
	}

	ss, _ := strconv.Atoi(a + a)
	sss, _ := strconv.Atoi(a + a + a)

	return s + ss + sss
}
func main() {

	val := myFunc("5")
	fmt.Println(val)

}
