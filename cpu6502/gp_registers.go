package cpu6502

type GP8Register struct {
	data uint8
}

func NewGP8() GP8 {
	return &GP8Register{
		data: 0,
	}
}

func (r *GP8Register) Get() uint8 {
	return r.data
}

func (r *GP8Register) Set(data uint8) {
	r.data = data
}
