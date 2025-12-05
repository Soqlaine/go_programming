package main

import "fmt"

func searchItem(a []string, b string) bool {

	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false

}
func main() {

	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")
	fmt.Println(result)

	result = searchItem(animals, "pig")
	fmt.Println(result)
}
