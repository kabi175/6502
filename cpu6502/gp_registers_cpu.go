package cpu6502

import (
	"github.com/kabi175/6502/model"
	"github.com/kabi175/6502/util"
)

type GP8Register struct {
	data uint8

	*util.Observer
}

func NewGP8() model.GP8 {
	return &GP8Register{
		data:     0,
		Observer: util.NewObserver(),
	}
}

func (r *GP8Register) Get() uint8 {
	return r.data
}

func (r *GP8Register) Set(data uint8) {
	r.data = data
}
