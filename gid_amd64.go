package gid

import "unsafe"

func Get() int64 {
	gptr := getg()
	return (*g)(gptr).goid
}

func getg() unsafe.Pointer
