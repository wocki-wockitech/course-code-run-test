package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Текущая версия: %s\n", runtime.Version())
	fmt.Println("Попробуйте переключить — доступны только 1.21, 1.22, 1.23")
}
