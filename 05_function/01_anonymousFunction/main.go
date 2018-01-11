package main

import "fmt"

func main() {
	x := 0
	// assign a function to variable increment
	// used to put other function into main function
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
