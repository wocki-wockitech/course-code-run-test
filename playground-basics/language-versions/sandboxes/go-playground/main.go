package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello from Go %s!\n", runtime.Version())
}
