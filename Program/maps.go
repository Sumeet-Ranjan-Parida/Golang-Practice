package main

import "fmt"

func main() {
	emails := map[string]string{"sumeet": "sumeet@abc.com", "sashank": "sashank@abc.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))

	delete(emails, "sashank")
	fmt.Println(emails)
}
