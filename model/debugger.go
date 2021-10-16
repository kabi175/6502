package model

type Debugger interface {
	Wait(uint16)
	IsEnd(uint16) bool
}
