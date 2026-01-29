package main

import "fmt"

type money float64

func (m money) printStr() string {
	return fmt.Sprintf("%.2f\n", m)
}

func main() {

	printmoney := []money{5.063, 9.2, 7.9111}
	//printmoney.print()

	for _, v := range printmoney {
		fmt.Println(v.printStr())
	}
}
