package caps

/*
#include <sys/capability.h>
#cgo LDFLAGS: -lcap
*/
import "C"

// GetProc() returns a capability set reflecting the capabilities of the
// calling process.
func GetProc() (*Cap, error) {
	c_cap, err := C.cap_get_proc()
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}

// SetProc() sets the capabilities of the calling process.
//
// If any flag is set for any capability not currently permitted for the
// calling process, the function will fail, and the capability state of the
// process will remain unchanged.
func (c Cap) SetProc() error {
	r, err := C.cap_set_proc(c.c)
	return _err(r, err)
}

// GetPid() returns a capability set reflecting the capabilities of the process
// indicated by pid.
//
// This information can also be obtained from the /proc/<pid>/status file.
func GetPid(pid int) (*Cap, error) {
	c_cap, err := C.cap_get_pid(C.int(pid))
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}
