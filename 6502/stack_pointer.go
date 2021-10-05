package cpu6502

type SP8 interface {
	Set(uint8)
	Get() uint8
	Reset()
}

type sp struct {
	pointer uint8
}

func NewSP8() SP8 {
	return &sp{
		pointer: 0,
	}
}

func (this *sp) Set(next uint8) {
	this.pointer = next
}

func (this *sp) Get() uint8 {
	return this.pointer
}

func (this *sp) Reset() {
	var initialState uint8 = 0
	this.Set(initialState)
}
