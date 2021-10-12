package cpu6502

type Bus16 interface {
	Read(uint16) uint8
	Write(uint16, uint8)
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

type Opcode interface {
	Execute(*Cpu6502) uint8
}

type FlagRegister interface {
	Set(int, bool)
	Get(int) bool
	Byte() uint8
	SetByte(uint8)
	Reset()
}

type Debugger interface {
	Render(*Cpu6502)
}
