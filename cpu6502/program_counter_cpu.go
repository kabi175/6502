package cpu6502

import (
	"github.com/kabi175/6502/model"
	"github.com/kabi175/6502/util"
)

type pc struct {
	counter uint16

	*util.Observer
}

func NewPC() model.PC16 {
	return &pc{
		counter:  0,
		Observer: util.NewObserver(),
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
