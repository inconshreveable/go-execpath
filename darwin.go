package execpath

// +build darwin

import (
	"fmt"
)

/*
#cgo CFLAGS: -mmacosx-version-min=10.6 -D__MAC_OS_X_VERSION_MAX_ALLOWED=1060

#include <mach-o/dyld.h>
#include <string.h>
*/
import "C"

func GetNative() (string, error) {
	var buflen C.uint32_t = C.uint32_t(maxPathSize)
	buf := make([]C.char, buflen)

	ret := C._NSGetExecutablePath(&buf[0], &buflen)
	if ret == -1 {
		// buflen wasn't large enough, _NSGetExecutablePath set it to the necessary size
		// so recreate the buffer and try again
		buf = make([]C.char, buflen)
		ret = C._NSGetExecutablePath(&buf[0], &buflen)
		if ret == -1 {
			// this should never happen
			return "", fmt.Errorf("_NSGetExecutable failed to get the executable path")
		}
	}
	pathlen := C.strnlen(&buf[0], C.size_t(buflen))
	return C.GoStringN(&buf[0], C.int(pathlen)), nil

}
