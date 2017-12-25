package stringutil

import "fmt"

// Visible fucntion will be exported
func Visible() {
	fmt.Println("This function is visible.")
	invisible()
}
