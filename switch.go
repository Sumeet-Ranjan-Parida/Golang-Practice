package main

import "fmt"

func main() {

	color := "red"

	switch color {

	case "red":
		fmt.Println("The color is red")

	case "blue":
		fmt.Println("The color is blue")
	default:
		fmt.Println("The color is not red or blue")
	}
}
