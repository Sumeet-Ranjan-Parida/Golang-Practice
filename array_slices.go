package main

import "fmt"

func main() {
	arr := [2]int{4, 5}

	array := [...]int{1, 2, 3, 4, 5}

	var matrix [3][3]int

	matrix[0] = [3]int{1, 1, 1}
	matrix[1] = [3]int{2, 2, 2}
	matrix[2] = [3]int{3, 3, 3}

	fmt.Println(arr)

	fmt.Println(array)

	fmt.Println(matrix)

	slice := []string{"a", "b"}

	fmt.Println(slice)

	m := make([]int, 5)

	fmt.Println(m)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
	b := a[:]
	c := a[2:]
	d := a[:8]
	e := a[4:8]
	a = append(a, 11)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(a)
}
