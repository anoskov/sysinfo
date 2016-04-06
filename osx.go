package sysinfo

import (
	"C"
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

	if err := sysctlbyname("kern.boottime", &tv); err != nil {
		return err
	}

	self.Length = time.Since(time.Unix(tv.Unix())).Seconds()

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
