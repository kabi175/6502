package cpu

import "github.com/kabi175/6502/model"

type CPU interface {
	Execute() // Execute method runs opcodes
	Quit()    // Quit method stops the cpu
	Reset()   // reset registers
}

type CPU6502 struct {
	PC uint16
	AC uint8
	X  uint8
	Y  uint8
	SR uint8
	SP uint8

	bus model.Bus16
}

var _ CPU = &CPU6502{}

func NewCPU(bus model.Bus16) CPU {
	return &CPU6502{bus: bus}
}

// Execute implements CPU.
func (c *CPU6502) Execute() {
	panic("unimplemented")
}

// read data from bus at PC & incre PC
func (c *CPU6502) FetchData() uint8 {
	data := c.Read(c.PC)
	c.PC++
	return data
}

// Quit implements CPU.
func (c *CPU6502) Quit() {
	panic("unimplemented")
}

// Reset implements CPU.
func (c *CPU6502) Reset() {
	c.AC = 0
	c.X = 0
	c.Y = 0
	c.SP = 0
	c.SR = 0
}

func (c *CPU6502) Read(add uint16) uint8 {
	return c.bus.Read(add)
}

func (c *CPU6502) Write(add uint16, data uint8) {
	c.bus.Write(add, data)
}
