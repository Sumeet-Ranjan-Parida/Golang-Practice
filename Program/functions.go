package main

import "fmt"

func greetings(name string) string {
	return "Hello, " + name
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greetings("Sumeet"))
	fmt.Println(add(2, 2))
}
