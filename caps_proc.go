package caps

/*
#include <sys/capability.h>
#cgo LDFLAGS: -lcap
*/
import "C"

func GetProc() (*Cap, error) {
	c_cap, err := C.cap_get_proc()
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}

func SetProc(c Cap) error {
	r, err := C.cap_set_proc(c.c)
	return _err(r, err)
}

func GetPid(pid int) (*Cap, error) {
	c_cap, err := C.cap_get_pid(C.int(pid))
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}
