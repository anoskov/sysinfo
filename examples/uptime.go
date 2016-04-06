package main

import (
	"github.com/anoskov/sysinfo"
	"fmt"
)

func main() {
	uptime := sysinfo.Uptime{}
	uptime.Get()

	fmt.Printf("Node Uptime: %s\n:", uptime.Format())
}
