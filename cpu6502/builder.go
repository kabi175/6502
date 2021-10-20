package cpu6502

import (
	"github.com/kabi175/6502/model"
)

//	Constructor function to create CPU6502
func NewCPU(bus model.Bus16, deb model.Debugger) *Cpu6502 {
	return &Cpu6502{
		PC:   NewPC(),
		SP:   NewSP8(),
		Flag: NewFlagRegister(),
		A:    NewGP8(),
		X:    NewGP8(),
		Y:    NewGP8(),
		Bus:  bus,
		deb:  deb,
	}
}
