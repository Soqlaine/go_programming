package main

import (
	"log"
	"os"
)

func main() {

	// opening the file in write-only mode if the file exists and then it truncates the file.
	// if the file doesn't exist it creates the file with 0644 permissions

	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	// error handling

	if err != nil {
		log.Fatal(err)
	}

	// defer closing the file
	defer file.Close()

	// WRITING BYTES TO FILE

	byteSlice := []byte("I learn golang! 传")

	bytesWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(bytesWritten)

	// WRITING BYTES TO FILE USING ioutil.WriteFile()

	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.

	bs := []byte("Go Programming is cool")

	err = os.WriteFile("c.txt", bs, 0644)

	// error handling

	if err != nil {
		log.Fatal(err)
	}
}
