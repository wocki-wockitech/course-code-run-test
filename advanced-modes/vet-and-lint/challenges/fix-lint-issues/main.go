package main

import (
	"fmt"
	"os"
)

func main() {
	name := "world"

	// Issue 1: wrong format verb
	fmt.Printf("Hello %d\n", name)

	// Issue 2: error ignored
	os.Open("/tmp/test.txt")

	// Issue 3: variable shadow
	x := 10
	if true {
		x := 20
		fmt.Println(x)
	}
	fmt.Println(x)
}
