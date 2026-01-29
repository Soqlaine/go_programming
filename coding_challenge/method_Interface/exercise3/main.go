package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) vat() float64 {
	return b.price * 0.09
}

// func (b book) discount() {
// 	(b).price = (b).price * 0.9

// }

//same as above
func (b *book) discount(p float64) {
	(*b).price = (*b).price * p

}

func main() {

	myBook := book{"Trial Book", 300}

	vat := myBook.vat()
	fmt.Printf("Vat:%v\n", vat)

	myBook.discount(0.9)
	fmt.Printf("%#v\n", myBook)

}
