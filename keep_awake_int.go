package main

type KeepAwake interface {
	Start() error
	Stop() error
}
