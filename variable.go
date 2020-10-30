package main

import "fmt"

func main() {
	name := "Sumeet"
	age := 22
	isCool := true
	float := 12.34

	fmt.Printf("%q this is a %T type\n", name, name)
	fmt.Printf("%v this is a %T type\n", age, age)
	fmt.Printf("%v this is a %T type\n", isCool, isCool)
	fmt.Printf("%v this is a %T type", float, float)

}
