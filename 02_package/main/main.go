package main

import (
	"fmt"

	"github.com/GoTutorial/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.ExportedVar)
	stringutil.Visible()
}
