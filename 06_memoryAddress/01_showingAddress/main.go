package main

import "fmt"

func main() {
	a := 10
	fmt.Println("Value of a is", a)
	fmt.Println("Address of a is", &a)
	fmt.Printf("%d\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 20
	fmt.Println(a)
}
