package main

import (
	execpath ".."
	"fmt"
	"reflect"
	"runtime"
)

type GetMethodT func() (string, error)

func main() {
	testMethod := func(getMethod GetMethodT) {
		name := runtime.FuncForPC(reflect.ValueOf(getMethod).Pointer()).Name()
		if path, err := getMethod(); err != nil {
			fmt.Printf("%s(): [ERROR] %v\n", name, err)
		} else {
			fmt.Printf("%s(): %s\n", name, path)
		}
	}

	methods := []GetMethodT{execpath.GetNative, execpath.GetArg0, execpath.GetPath, execpath.Get}
	for _, method := range methods {
		testMethod(method)
	}
}
