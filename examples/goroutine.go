package main

import (
	"github.com/anoskov/sysinfo"
	"fmt"
)

func main() {
	goroutine := sysinfo.Goroutine{}
	goroutine.Get()

	fmt.Printf("Goroutine count: %d\n", goroutine.Count)
}
