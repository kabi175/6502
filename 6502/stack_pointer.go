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

func (s *sp) Set(next uint8) {
	s.pointer = next
}

func (s *sp) Get() uint8 {
	return s.pointer
}

func (s *sp) Reset() {
	var initialState uint8 = 0
	s.Set(initialState)
}
