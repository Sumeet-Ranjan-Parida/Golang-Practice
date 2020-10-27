package main

import "fmt"

type Fruits struct {
	number int
	name   string
	Fruits []string
}

func main() {
	a := Fruits{
		number: 3,
		name:   "Red Fruits",
		Fruits: []string{
			"Apple",
			"Cherries",
			"Strawberries",
		},
	}
	fmt.Println(a)
	fmt.Println(a.Fruits)

	b := Fruits{
		number: 3,
		name:   "Green Fruits",
		Fruits: []string{
			"Grapes",
			"Watermelon",
			"Kiwi",
		},
	}

	fmt.Println(b)
	fmt.Println(b.Fruits)

}
