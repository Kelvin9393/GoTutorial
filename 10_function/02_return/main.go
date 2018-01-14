package main

import "fmt"

func main() {
	fmt.Println(greet("Kelvin ", "Ling"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
