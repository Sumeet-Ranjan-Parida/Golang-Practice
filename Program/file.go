package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, error := os.Create("sample.txt")
	if error != nil {
		log.Fatal(error)
	}
	file.WriteString("Hello, world!!!")
	file.Close()
	stream, error := ioutil.ReadFile("sample.txt")
	if error != nil {
		log.Fatal(error)
	}
	view := string(stream)
	fmt.Println(view)
}
