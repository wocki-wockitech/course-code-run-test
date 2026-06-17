package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Эта песочница зафиксирована на Go %s\n", runtime.Version())
	fmt.Println("Переключить версию нельзя — селектор скрыт.")
}
