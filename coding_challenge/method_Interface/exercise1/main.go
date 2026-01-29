package main

import "fmt"

type money float64

func (m money) print() {
	fmt.Printf("%.2f\n", m)
}

func main() {
	printmoney := []money{5.063, 9.2, 7.9111}
	//printmoney.print()

	for _, v := range printmoney {
		v.print()
	}
}
