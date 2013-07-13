package execpath

// +build linux

import (
	"os"
)

func GetNative() (string, error) {
	return os.Readlink("/proc/self/exe")
}
