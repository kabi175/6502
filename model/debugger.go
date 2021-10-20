package model

// A Debugger is capable of pause and end the CPU Process
type Debugger interface {
	Wait(uint16)
	IsEnd(uint16) bool
}
