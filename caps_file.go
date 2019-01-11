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

// GetFile reads a capability state from the given file.
//
// The effects of reading the capability state from any file other than a
// regular file is undefined.
func GetFile(f *os.File) (*Cap, error) {
	c_cap, err := C.cap_get_fd(C.int(f.Fd()))
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}

// GetFilePath reads a capability state from the given file.
//
// The effects of reading the capability state from any file other than a
// regular file is undefined.
func GetFilePath(path string) (*Cap, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	c_cap, err := C.cap_get_file(cPath)
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}

// SetFile set the values for all capability flags for all capabilities for the
// file with the given capability state.
//
// For this functions to succeed, the calling process must have the effective
// capability, CAP_SETFCAP, enabled and either the effective user ID of the
// process must match the file owner or the calling process must have the
// CAP_FOWNER flag in its effective capability set. The effects of writing the
// capability state to any file type other than a regular file are undefined.
func SetFile(f *os.File, c Cap) error {
	r, err := C.cap_set_fd(C.int(f.Fd()), c.c)
	return _err(r, err)
}

// SetFilePath set the values for all capability flags for all capabilities for
// the file with the given capability state.
//
// For this functions to succeed, the calling process must have the effective
// capability, CAP_SETFCAP, enabled and either the effective user ID of the
// process must match the file owner or the calling process must have the
// CAP_FOWNER flag in its effective capability set. The effects of writing the
// capability state to any file type other than a regular file are undefined.
func (c Cap) SetFilePath(path string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	r, err := C.cap_set_file(cPath, c.c)

	return _err(r, err)
}
