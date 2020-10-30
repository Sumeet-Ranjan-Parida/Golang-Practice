package main

import "fmt"

func main() {
	color := "red"

	if color == "green" {
		fmt.Println("The color is green not red")
	} else if color == "red" {
		fmt.Println("The color is red")
	} else {
		fmt.Println("The color is neither green nor red")
	}
}
