package opcode

import cpu "github.com/kabi175/6502/cpu6502"

type opcode struct {
	Cycle uint8
	mode  string
	Mode  func(*cpu.Cpu6502) uint8
	Ins   func(*cpu.Cpu6502) uint8
}

func NewOpcode() cpu.Opcode {
	return &opcode{}
}

func (i *opcode) Execute(c *cpu.Cpu6502) uint8 {
	cycles := i.Cycle
	cycles += i.Mode(c)
	cycles += i.Ins(c)
	return cycles
}
