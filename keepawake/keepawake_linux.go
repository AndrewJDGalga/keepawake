package keepawake

import "github.com/godbus/dbus/v5"

type linuxAwake struct {
	conn *dbus.Conn
	fd   dbus.UnixFD
}

func New() KeepAwake {
	return &linuxAwake{}
}

func (w *linuxAwake) Start() error {

	return nil
}

func (w *linuxAwake) Stop() error {

	return nil
}
