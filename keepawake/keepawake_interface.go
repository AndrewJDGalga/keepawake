package keepawake

type KeepAwake interface {
	Start() error
	Stop() error
}
