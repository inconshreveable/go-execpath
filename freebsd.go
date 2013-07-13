package execpath

// +build freebsd

import (
	"os"
)

func GetNative() (path string, err error) {
	return os.Readlink("/proc/curproc/file")
}
