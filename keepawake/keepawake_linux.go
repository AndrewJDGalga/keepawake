package keepawake

import (
	"syscall"

	"github.com/godbus/dbus/v5"
)

type linuxAwake struct {
	sysConn  *dbus.Conn
	fd       dbus.UnixFD
	sessConn *dbus.Conn
	cookie   uint32
}

func New() KeepAwake {
	return &linuxAwake{}
}

func (l *linuxAwake) Start() error {
	sysConn, err := dbus.ConnectSystemBus()
	if err != nil {
		return err
	}
	l.sysConn = sysConn

	sysObj := sysConn.Object("org.freedesktop.login1", "/org/freedesktop/login1")
	sysCall := sysObj.Call(
		//needs what, who, why, mode
		"org.freedesktop.login1.Manager.Inhibit",
		0,
		"idle:sleep",
		"keepawake",
		"User requested keepawake",
		"block",
	)
	if sysCall.Err != nil {
		return sysCall.Err
	}

	if err := sysCall.Store(&l.fd); err != nil {
		return err
	}

	sessConn, err := dbus.ConnectSessionBus()
	if err != nil {
		return err
	}
	l.sessConn = sessConn

	sessObj := sessConn.Object("org.freedesktop.ScreenSaver", "/org/freedesktop/ScreenSaver")
	sessCall := sessObj.Call(
		"org.freedesktop.ScreenSaver.Inhibit",
		0,
		"keepawake",
		"User requested keepawake",
	)
	if sessCall.Err != nil {
		return sessCall.Err
	}

	return sessCall.Store(&l.cookie)
}

func (l *linuxAwake) Stop() error {
	if err := syscall.Close(int(l.fd)); err != nil {
		return err
	}
	if err := l.sysConn.Close(); err != nil {
		return err
	}

	sessObj := l.sessConn.Object("org.freedesktop.ScreenSaver", "/org/freedesktop/ScreenSaver")
	if call := sessObj.Call("org.freedesktop.ScreenSaver.UnInhibit", 0, l.cookie); call.Err != nil {
		return call.Err
	}

	return l.sessConn.Close()
}
