package main

//https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-mouseinput
type mouseInput struct {
	dx, dy      int32  //LONG
	mouseData   uint32 //DWORD
	dwFlags     uint32
	time        uint32
	dwExtraInfo uintptr
}
