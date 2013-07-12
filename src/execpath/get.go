/*
   execpath provides functions for determining the absolute path to the
   executable file of the running program.
*/
package execpath

var maxPathSize = 1024

func Get() (path string, err error) {
	if path, err = GetNative(); err == nil {
		return
	}

	if path, err = GetArg0(); err == nil {
		return
	}

	if path, err = GetPath(); err == nil {
		return
	}

	return
}
