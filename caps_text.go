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

func FromText(text string) (*Cap, error) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	c_cap, err := C.cap_from_text(cText)
	if c_cap == nil {
		return nil, err
	}

	return create(c_cap), nil
}

func (c Cap) ToText() (string, error) {
	text, err := C.cap_to_text(c.c, nil)
	defer C.cap_free(unsafe.Pointer(text))

	if text == nil {
		return "", err
	}

	return C.GoString(text), nil
}

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

func ToName(value CapValue) string {
	var s = C.cap_to_name(C.cap_value_t(value))
	C.cap_free(unsafe.Pointer(s))

	return C.GoString(s)
}
