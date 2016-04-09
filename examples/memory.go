package main

import (
	"fmt"
	"github.com/anoskov/sysinfo"
	"os"
)

func format(val uint64) uint64 {
	return val / 1024
}

func main() {
	mem := sysinfo.RAM{}

	mem.Get()

	fmt.Fprintf(os.Stdout, "%18s %10s %10s\n",
		"total", "used", "free")

	fmt.Fprintf(os.Stdout, "Mem:    %10d %10d %10d\n",
		format(mem.Total), format(mem.Used), format(mem.Free))

	fmt.Fprintf(os.Stdout, "-/+ buffers/cache: %10d %10d\n",
		format(mem.ActualUsed), format(mem.ActualFree))
}
