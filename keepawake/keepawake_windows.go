package keepawake

import "golang.org/x/sys/windows"

const (
	// https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-setthreadexecutionstate
	SYS_CONTINOUS_STATE = 0x80000000
	SYS_WORKING_STATE   = 0x00000001
	SYS_DISPLAY_REQ     = 0x00000002
)

var (
	winKernal        = windows.NewLazySystemDLL("kernel32.dll")
	pThreadExecState = winKernal.NewProc("SetThreadExecutionState")
)

type windowsAwake struct{}

func New() KeepAwake {
	return &windowsAwake{}
}

func (w *windowsAwake) Start() error {
	prevState, _, err := pThreadExecState.Call(
		SYS_CONTINOUS_STATE |
			SYS_WORKING_STATE |
			SYS_DISPLAY_REQ,
	)
	if prevState == 0 {
		return err
	}
	return nil
}

func (w *windowsAwake) Stop() error {
	//somehow the reset state is setting simply continous
	prevState, _, err := pThreadExecState.Call(
		SYS_CONTINOUS_STATE,
	)
	if prevState == 0 {
		return err
	}
	return nil
}
