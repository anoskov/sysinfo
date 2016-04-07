package main

import (
	"github.com/anoskov/sysinfo"
	"fmt"
)

func main() {
	cpu := sysinfo.CPU{}
	cpu.Get()

	fmt.Printf("CPU count: %d\n", cpu.Count)
}
