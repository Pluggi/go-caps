package caps

/*
#include <sys/capability.h>
#cgo LDFLAGS: -lcap
*/
import "C"
import (
	"errors"
)

// Clear() initializes the capability state in working storage so that all
// capability flags are cleared.
func (c Cap) Clear() error {
	r, err := C.cap_clear(c.c)
	return _err(r, err)
}

// ClearFlag() clears all of the capabilities of the specified capability flag.
func (c Cap) ClearFlag(flag CapFlag) error {
	r, err := C.cap_clear_flag(c.c, C.cap_flag_t(flag))
	return _err(r, err)
}

var (
	ErrCapNotEqual = errors.New("Capabilities not equal")
)

// Compare() compares two full capability sets and returns nil if the two
// capability sets are identical.
//
// A difference between the two sets returns ErrCapNotEqual.
func Compare(a, b Cap) error {
	r, err := C.cap_compare(a.c, b.c)
	if r < 0 {
		return err
	}
	if r > 0 {
		return ErrCapNotEqual
	}

	return nil
}

// GetFlag() returns the current value of the capability flag.
func (c Cap) GetFlag(cap_value CapValue, flag CapFlag) (CapFlagValue, error) {
	value := C.cap_flag_value_t(0)

	r, err := C.cap_get_flag(c.c, C.int(cap_value), C.cap_flag_t(flag), &value)
	if r < 0 {
		return -1, err
	}

	return CapFlagValue(value), nil
}

// SetFlag() sets the flag of each capability in the slice caps to the
// CapFlagValue value.
func (c Cap) SetFlag(flag CapFlag, caps []CapValue, value CapFlagValue) error {
	capsint_ncap := make([]C.int, len(caps))
	for i := 0; i < len(caps); i++ {
		capsint_ncap[i] = C.int(caps[i])
	}

	r, err := C.cap_set_flag(
		c.c,
		C.cap_flag_t(flag),
		C.int(len(caps)),
		&capsint_ncap[0],
		C.cap_flag_value_t(value),
	)

	return _err(r, err)
}
