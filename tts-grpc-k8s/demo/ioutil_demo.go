package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	tempFile()
}

func tempFile() {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file name :%s", file.Name())
}
