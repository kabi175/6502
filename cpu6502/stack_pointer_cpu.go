package cpu6502

import (
	"github.com/kabi175/6502/model"
	"github.com/kabi175/6502/util"
)

type sp struct {
	pointer uint8
	*util.Observer
}

func NewSP8() model.SP8 {
	return &sp{
		pointer:  0xFF,
		Observer: util.NewObserver(),
	}
}

func (s *sp) Set(next uint16) {
	s.pointer = uint8(next & 0xff)
}

func (s *sp) Get() uint16 {
	return uint16(s.pointer & 0xff)
}

func (s *sp) Reset() {
	initialState := uint16(0)
	s.Set(initialState)
}

func (s *sp) Increment() {
	s.Set(s.Get() + 1)
}

func (s *sp) Decrement() {
	s.Set(s.Get() - 1)
}
