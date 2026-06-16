package main

import (
	"fmt"
	"os"
)

func main() {
	name := "world"

	// Fixed: correct format verb
	fmt.Printf("Hello %s\n", name)

	// Fixed: handle error
	f, err := os.Open("/tmp/test.txt")
	if err != nil {
		fmt.Println("file not found, that's ok")
	} else {
		f.Close()
	}

	// Fixed: no shadow — reuse outer x
	x := 10
	if true {
		x = 20
		fmt.Println(x)
	}
	fmt.Println(x)
}
