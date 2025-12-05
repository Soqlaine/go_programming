package main

import (
	"log"
	"os"
)

func main() {

	bs := []byte("The Go gopher is an iconic mascot!")
	err := os.WriteFile("info.txt", bs, 0644)

	if err != nil {
		log.Fatal(err)
	}
}
