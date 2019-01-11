package caps

/*
#include <sys/capability.h>
#cgo LDFLAGS: -lcap
*/
import "C"
import (
	"runtime"
	"unsafe"
)

// Equivalent of cap_value_t
type CapValue int

// Generated using
//    awk '$1 == "#define" &&                                    \
//         $2 ~ /^CAP_\w+$/ &&                                   \
//         $2 != "CAP_LAST_CAP"                                  \
//         { printf("%-20s%s= CapValue(C.%s)\n", $2, " ", $2) }' \
//         /usr/include/linux/capability.h
const (
	CAP_CHOWN            = CapValue(C.CAP_CHOWN)
	CAP_DAC_OVERRIDE     = CapValue(C.CAP_DAC_OVERRIDE)
	CAP_DAC_READ_SEARCH  = CapValue(C.CAP_DAC_READ_SEARCH)
	CAP_FOWNER           = CapValue(C.CAP_FOWNER)
	CAP_FSETID           = CapValue(C.CAP_FSETID)
	CAP_KILL             = CapValue(C.CAP_KILL)
	CAP_SETGID           = CapValue(C.CAP_SETGID)
	CAP_SETUID           = CapValue(C.CAP_SETUID)
	CAP_SETPCAP          = CapValue(C.CAP_SETPCAP)
	CAP_LINUX_IMMUTABLE  = CapValue(C.CAP_LINUX_IMMUTABLE)
	CAP_NET_BIND_SERVICE = CapValue(C.CAP_NET_BIND_SERVICE)
	CAP_NET_BROADCAST    = CapValue(C.CAP_NET_BROADCAST)
	CAP_NET_ADMIN        = CapValue(C.CAP_NET_ADMIN)
	CAP_NET_RAW          = CapValue(C.CAP_NET_RAW)
	CAP_IPC_LOCK         = CapValue(C.CAP_IPC_LOCK)
	CAP_IPC_OWNER        = CapValue(C.CAP_IPC_OWNER)
	CAP_SYS_MODULE       = CapValue(C.CAP_SYS_MODULE)
	CAP_SYS_RAWIO        = CapValue(C.CAP_SYS_RAWIO)
	CAP_SYS_CHROOT       = CapValue(C.CAP_SYS_CHROOT)
	CAP_SYS_PTRACE       = CapValue(C.CAP_SYS_PTRACE)
	CAP_SYS_PACCT        = CapValue(C.CAP_SYS_PACCT)
	CAP_SYS_ADMIN        = CapValue(C.CAP_SYS_ADMIN)
	CAP_SYS_BOOT         = CapValue(C.CAP_SYS_BOOT)
	CAP_SYS_NICE         = CapValue(C.CAP_SYS_NICE)
	CAP_SYS_RESOURCE     = CapValue(C.CAP_SYS_RESOURCE)
	CAP_SYS_TIME         = CapValue(C.CAP_SYS_TIME)
	CAP_SYS_TTY_CONFIG   = CapValue(C.CAP_SYS_TTY_CONFIG)
	CAP_MKNOD            = CapValue(C.CAP_MKNOD)
	CAP_LEASE            = CapValue(C.CAP_LEASE)
	CAP_AUDIT_WRITE      = CapValue(C.CAP_AUDIT_WRITE)
	CAP_AUDIT_CONTROL    = CapValue(C.CAP_AUDIT_CONTROL)
	CAP_SETFCAP          = CapValue(C.CAP_SETFCAP)
	CAP_MAC_OVERRIDE     = CapValue(C.CAP_MAC_OVERRIDE)
	CAP_MAC_ADMIN        = CapValue(C.CAP_MAC_ADMIN)
	CAP_SYSLOG           = CapValue(C.CAP_SYSLOG)
	CAP_WAKE_ALARM       = CapValue(C.CAP_WAKE_ALARM)
	CAP_BLOCK_SUSPEND    = CapValue(C.CAP_BLOCK_SUSPEND)
	CAP_AUDIT_READ       = CapValue(C.CAP_AUDIT_READ)
)

// Equivalent of cap_flag_t
type CapFlag int

const (
	CAP_EFFECTIVE   = CapFlag(C.CAP_EFFECTIVE)
	CAP_PERMITTED   = CapFlag(C.CAP_PERMITTED)
	CAP_INHERITABLE = CapFlag(C.CAP_INHERITABLE)
)

// Equivalent of cap_flag_value_t
type CapFlagValue int

const (
	CAP_CLEAR = CapFlagValue(C.CAP_CLEAR)
	CAP_SET   = CapFlagValue(C.CAP_SET)
)

// Wrapping structure for cap_t
type Cap struct {
	c C.cap_t
}

// NewCap() allocates a new Cap structure.
func NewCap() Cap {
	c := Cap{
		C.cap_init(),
	}
	runtime.SetFinalizer(c, freeCap)

	return c
}

// create() allocates a new Cap structure using c_cap
func create(c_cap C.cap_t) *Cap {
	c := &Cap{c_cap}
	runtime.SetFinalizer(c, freeCap)

	return c
}

// freeCap frees a C.cap_t object using cap_free
func freeCap(c *Cap) {
	C.cap_free(unsafe.Pointer(&c.c))
}

// _err takes a return code and errno as parameters and returns an error if
// there was an error, else nil
//
// Example:
//   C:
//     int r = cap_set_fd(fd, c)
//     if (r < 0) {
//       return errno
//     }
//     return r
//
//   Go:
//     r, err := C.cap_set_fd(fd, c)
//     return _err(r, err)
func _err(r C.int, err error) error {
	if r < 0 {
		return err
	}

	return nil
}
