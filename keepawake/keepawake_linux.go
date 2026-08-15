package keepawake

import "github.com/godbus/dbus/v5"

type linuxAwake struct {
	conn *dbus.Conn
	fd   dbus.UnixFD
}

func New() KeepAwake {
	return &linuxAwake{}
}

func (l *linuxAwake) Start() error {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		return err
	}
	l.conn = conn

	obj := conn.Object("org.freedesktop.login1", "org/freedesktop/login1")
	call := obj.Call(
		//needs what, who, why, mode
		"org.freedesktop.login1.Manager.Inhibit",
		0,
		"idle:sleep",
		"keepawake",
		"USer requested keepawake",
		"block",
	)
	if call.Err != nil {
		return call.Err
	}

	return call.Store(&l.fd)
}

func (w *linuxAwake) Stop() error {

	return nil
}
