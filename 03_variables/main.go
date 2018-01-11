package main

import "fmt"

var x int = 5

func main() {
    a := 10
    b := "golang"
    c := 4.32
    d := true

    fmt.Printf("%v\n", a)
    fmt.Printf("%v\n", b)
    fmt.Printf("%v\n", c)
    fmt.Printf("%v\n", d)

    fmt.Println(x)
}
