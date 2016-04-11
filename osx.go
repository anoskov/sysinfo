package sysinfo

/*
#include <stdlib.h>
#include <sys/sysctl.h>
#include <sys/mount.h>
#include <mach/mach_init.h>
#include <mach/mach_host.h>
#include <mach/host_info.h>
#include <libproc.h>
#include <mach/processor_info.h>
#include <mach/vm_map.h>
*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"syscall"
	"time"
	"unsafe"
	"fmt"
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

func (self *RAM) Get() error {
	var vmstat C.vm_statistics_data_t

	if err := sysctlByName("hw.memsize", &self.Total); err != nil {
		return err
	}

	if err := vm_info(&vmstat); err != nil {
		return err
	}

	kern := uint64(vmstat.inactive_count) << 12
	self.Free = uint64(vmstat.free_count) << 12

	self.Used = self.Total - self.Free
	self.ActualFree = self.Free + kern
	self.ActualUsed = self.Used - kern

	return nil
}

type xsw_usage struct {
	Total, Avail, Used uint64
}

func (self *Swap) Get() error {
	sw_usage := xsw_usage{}

	if err := sysctlByName("vm.swapusage", &sw_usage); err != nil {
		return err
	}

	self.Total = sw_usage.Total
	self.Used = sw_usage.Used
	self.Free = sw_usage.Avail

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

func vm_info(vmstat *C.vm_statistics_data_t) error {
	var count C.mach_msg_type_number_t = C.HOST_VM_INFO_COUNT

	status := C.host_statistics(
		C.host_t(C.mach_host_self()),
		C.HOST_VM_INFO,
		C.host_info_t(unsafe.Pointer(vmstat)),
		&count)

	if status != C.KERN_SUCCESS {
		return fmt.Errorf("host_statistics=%d", status)
	}

	return nil
}
