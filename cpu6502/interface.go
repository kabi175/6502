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
	Reset()
}
