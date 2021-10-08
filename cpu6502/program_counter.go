package cpu6502

type pc struct {
	counter uint16
}

func NewPC() PC16 {
	return &pc{
		counter: 0,
	}
}

func (p *pc) Set(next uint16) {
	p.counter = next
}

func (p *pc) Get() uint16 {
	return p.counter
}

func (p *pc) Reset() {
	var initialState uint16 = 0
	p.Set(initialState)
}

func (p *pc) Increment() {
	nextState := p.Get()
	p.Set(nextState)
}
