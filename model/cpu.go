package model

// An interface for Cpu
type Cpu interface {
	Execute(chan bool) // Execute method runs opcodes
	State() []byte     // Get Cpu State in JSON fromat
}

type GP8 interface {
	Get() uint8
	Set(uint8)
}

type PC16 interface {
	Set(uint16)
	Get() uint16
	Increment()
	Decrement()
	Reset()
}

type SP8 interface {
	Set(uint16)
	Get() uint16
	Increment()
	Decrement()
	Reset()
}

type FlagRegister interface {
	Set(int, bool)
	Get(int) bool
	Byte() uint8
	SetByte(uint8)
	Reset()
}
