package main

import (
	"fmt"
	"github.com/anoskov/sysinfo"
)

func main() {
	avload := sysinfo.AverageLoad{}
	avload.Get()

	fmt.Printf("Average Load: %.2f, %.2f, %.2f\n", avload.One, avload.Five, avload.Fifteen)
}
