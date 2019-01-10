package caps

/*
#include <stdlib.h>
#include <sys/capability.h>
#cgo LDFLAGS: -lcap
*/
import "C"
import (
	"os"
	"unsafe"
)

func GetFile(f *os.File) (*Cap, error) {
	c_cap, err := C.cap_get_fd(C.int(f.Fd()))
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}

func GetFilePath(path string) (*Cap, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	c_cap, err := C.cap_get_file(cPath)
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}

func SetFile(f *os.File, c Cap) error {
	r, err := C.cap_set_fd(C.int(f.Fd()), c.c)
	return _err(r, err)
}

func (c Cap) SetFilePath(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	r, err := C.cap_set_file(cPath, c.c)

	return _err(r, err)
}
