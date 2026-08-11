package main

import (
	"golang.org/x/sys/windows"
)

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-mouseinput
type MouseInput struct {
	Dx, Dy      int32  //LONG
	MouseData   uint32 //DWORD
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-input
type Input struct {
	Type uint32
	Mi   MouseInput
}

const MOUSE_EVENT_MOVE = 0x0001

var (
	userProc = windows.NewLazySystemDLL("user32.dll")
	pInput   = userProc.NewProc("SendInput")
)
