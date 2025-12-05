package main

import (
	"log"
	"os"
)

func main() {

	_, err := os.Stat("info.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}

	err = os.Rename("info.txt", "abba.txt")
	if err != nil {
		log.Fatal(err)
	}
}
