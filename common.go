package sysinfo

import "runtime"

func (self *CPU) Get() error {
	self.Count = runtime.NumCPU()

	return nil
}

func (self *Goroutine) Get() error {
	self.Count = runtime.NumGoroutine()

	return nil
}