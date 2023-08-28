package main

import (
	"fmt"
	"strings"

	pds "github.com/georgetsouvaltzis/pds"
)

func main() {
	fmt.Println(pds.Bark())
}

func WhenGrownUp(s string) string {
	return "When the puppy grows up it says: " + strings.ToUpper(s)
}
