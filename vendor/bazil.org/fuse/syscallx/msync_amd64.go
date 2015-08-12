// mksyscall.pl msync.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscallx

import "syscall"

import "unsafe"

// THIS FILE IS GENERATED BY THE COMMAND AT THE TOP; DO NOT EDIT

func Msync(b []byte, flags int) (err error) {
	var _p0 unsafe.Pointer
	if len(b) > 0 {
		_p0 = unsafe.Pointer(&b[0])
	} else {
		_p0 = unsafe.Pointer(&_zero)
	}
	_, _, e1 := syscall.Syscall(syscall.SYS_MSYNC, uintptr(_p0), uintptr(len(b)), uintptr(flags))
	if e1 != 0 {
		err = e1
	}
	return
}
