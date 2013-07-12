package execpath

// +build netbsd

import (
	"os"
)

func GetNative() (path string, err error) {
	return os.Readlink("/proc/curproc/exe")
}
