package sysinfo

/*
#include <stdlib.h>
*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"syscall"
	"time"
	"unsafe"
)

/*********************
**		    **
**     Methods      **
**		    **
**********************/

func (self *Uptime) Get() error {
	tv := syscall.Timeval{}

	if err := sysctlByName("kern.boottime", &tv); err != nil {
		return err
	}

	self.Duration = time.Since(time.Unix(tv.Unix())).Seconds()

	return nil
}

func (self *AverageLoad) Get() error {
	avg := []C.double{0, 0, 0}

	C.getloadavg(&avg[0], C.int(len(avg)))

	self.One = float64(avg[0])
	self.Five = float64(avg[1])
	self.Fifteen = float64(avg[2])

	return nil
}

/*********************
**		    **
**  Util Functions  **
**		    **
**********************/

// zsyscall_darwin_amd64.go

func sysctl(mib []C.int, old *byte, oldlen *uintptr, new *byte, newlen uintptr) (err error) {
	var ptr unsafe.Pointer
	ptr = unsafe.Pointer(&mib[0])
	_, _, e1 := syscall.Syscall6(syscall.SYS___SYSCTL, uintptr(ptr),
		uintptr(len(mib)),
		uintptr(unsafe.Pointer(old)), uintptr(unsafe.Pointer(oldlen)),
		uintptr(unsafe.Pointer(new)), uintptr(newlen))
	if e1 != 0 {
		err = e1
	}
	return
}

func sysctlByName(name string, data interface{}) (err error) {
	val, err := syscall.Sysctl(name)
	if err != nil {
		return err
	}

	buf := []byte(val)

	switch v := data.(type) {
	case *uint64:
		*v = *(*uint64)(unsafe.Pointer(&buf[0]))
		return
	}

	bbuf := bytes.NewBuffer([]byte(val))

	return binary.Read(bbuf, binary.LittleEndian, data)
}
