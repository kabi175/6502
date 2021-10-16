package cpu6502

import "github.com/kabi175/6502/model"

type pc struct {
	counter uint16
}

func NewPC() model.PC16 {
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
	initialState := uint16(0)
	p.Set(initialState)
}

func (p *pc) Increment() {
	nextState := p.Get() + 1
	p.Set(nextState)
}

func (p *pc) Decrement() {
	nextState := p.Get() - 1
	p.Set(nextState)
}
