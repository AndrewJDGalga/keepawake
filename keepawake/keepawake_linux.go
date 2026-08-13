package keepawake

type linuxAwake struct{}

func New() KeepAwake {
	return &linuxAwake{}
}

func (w *linuxAwake) Start() error {

	return nil
}

func (w *linuxAwake) Stop() error {

	return nil
}
