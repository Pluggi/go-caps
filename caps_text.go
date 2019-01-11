package caps

/*
#include <stdlib.h>
#include <sys/capability.h>
#cgo LDFLAGS: -lcap
*/
import "C"
import (
	"unsafe"
)

// FromText() returns a capability set reflecting the state represented by a
// human-readable capability set.
func FromText(text string) (*Cap, error) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	c_cap, err := C.cap_from_text(cText)
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}

// Returns a human-readable string of the capability set.
//
// Equivalent to cap_to_text(cap_t, ssize_t)
func (c Cap) String() (string, error) {
	text, err := C.cap_to_text(c.c, nil)
	defer C.cap_free(unsafe.Pointer(text))

	if text == nil {
		return "", err
	}

	return C.GoString(text), nil
}

// FromName() converts a text representation of a capability, such as
// "cap_chown", to a CapValue.
func FromName(name string) (CapValue, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	var value C.cap_value_t
	r, err := C.cap_from_name(cName, &value)
	if r < 0 {
		return -1, err
	}

	return CapValue(value), nil
}

// Converts a CapValue to a string
//
// Equivalent to cap_to_name(cap_value_t)
func (value CapValue) String() string {
	var s = C.cap_to_name(C.cap_value_t(value))
	C.cap_free(unsafe.Pointer(s))

	return C.GoString(s)
}
