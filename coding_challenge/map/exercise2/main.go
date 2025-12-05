package main

import "fmt"

func main() {
	// var m1 map[int]bool
	// m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	// fmt.Println(m2 == m3)

	mm2 := fmt.Sprintf("%s", m2)
	mm3 := fmt.Sprintf("%s", m3)

	if mm2 == mm3 {
		fmt.Println("It's equal")
	} else {
		fmt.Println("It's not equal")
	}

	var m1 map[int]bool

	// ERROR -> panic: assignment to entry in nil map
	// m1[5] = true

	f2 := map[int]int{3: 10, 4: 40}
	f3 := map[int]int{3: 10, 4: 40}

	// ERROR -> invalid operation: m2 == m3 (map can only be compared to nil)
	// fmt.Println(m2 == m3)

	_, _, _ = m1, f2, f3

}
