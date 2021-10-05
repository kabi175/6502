package cpu6502

type PC16 interface {
	Set(uint16)
	Get() uint16
	Reset()
}

type pc struct {
	counter uint16
}

func NewPC() PC16 {
	return &pc{
		counter: 0,
	}
}

func (this *pc) Set(next uint16) {
	this.counter = next
}

func (this *pc) Get() uint16 {
	return this.counter
}

func (this *pc) Reset() {
	var initialState uint16 = 0
	this.Set(initialState)
}
