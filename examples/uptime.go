package main

import (
	"fmt"
	"github.com/anoskov/sysinfo"
)

func main() {
	uptime := sysinfo.Uptime{}
	uptime.Get()

	fmt.Printf("Node Uptime: %s\n:", uptime.Format())
}
