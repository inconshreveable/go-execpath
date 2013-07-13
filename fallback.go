package execpath

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

func pathExists(p string) (bool, error) {
	_, err := os.Stat(p)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}

func makeAbsolute(p string) (fullPath string, err error) {
	if path.IsAbs(p) {
		// the shell gave us the absolute path, we're done
		fullPath = p
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		return
	}

	fullPath = path.Join(wd, p)
	return
}

// GetEnv() attempts to retrieve the executable path by examining the
// environment variable "_"
func GetEnv() (p string, err error) {
	p = os.Getenv("_")

	if p == "" {
		err = fmt.Errorf("Executable path not set in the environment")
	}

	return makeAbsolute(p)
}

// GetArg0() attempts to retrieve the executable path by examining the value
// of ARGV[0] and combining it with the working directory, if necessary
func GetArg0() (p string, err error) {
	p, err = makeAbsolute(os.Args[0])
	if err != nil {
		return
	}

	var exists bool
	if exists, err = pathExists(p); !exists {
		err = fmt.Errorf("Can't determine executable path from arg0")
	}

	return
}

// GetPath() attempts to retrieve the executable path by searching for the
// executable (os.Args[0]) in each directory in the PATH environment variable
func GetPath() (p string, err error) {
	executable := os.Args[0]
	envPath := os.Getenv("PATH")

	var sep string
	if runtime.GOOS == "windows" {
		if !strings.HasSuffix(executable, ".exe") {
			executable = executable + ".exe"
		}
		sep = ";"
	} else {
		sep = ":"
	}
	pathDirs := strings.Split(envPath, sep)

	for _, dir := range pathDirs {
		p = path.Join(dir, executable)
		var exists bool
		if exists, err = pathExists(p); err == nil && exists {
			return
		}
	}

	err = fmt.Errorf("Executable %s not found in path", executable)
	return
}
