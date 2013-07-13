package execpath

// +build plan9

import (
	"fmt"
)

func GetNative() (string, error) {
	return "", fmt.Errorf("GetNative() not implemented on plan9")
}
