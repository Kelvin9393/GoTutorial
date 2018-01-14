package main

import "fmt"

func main() {
	greet("Kelvin", "Ling")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
