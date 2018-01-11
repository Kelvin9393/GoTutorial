package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)

	{
		fmt.Println(x)
		y := "This variable y is inside a closure."
		fmt.Println(y)
	}

	// fmt.Println(y) // outside scope of y
}
