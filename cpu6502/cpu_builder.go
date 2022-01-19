package cpu6502

import (
	"github.com/kabi175/6502/model"
)

//	Constructor function to create CPU6502
func newCPU(bus model.Bus16) *CPU6502 {
	return &CPU6502{
		PC:   NewPC(),
		SP:   NewSP8(),
		Flag: NewFlagRegister(),
		A:    NewGP8(),
		X:    NewGP8(),
		Y:    NewGP8(),

		Bus: bus,
	}
}
