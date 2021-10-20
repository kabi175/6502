package opcode

import (
	cpu "github.com/kabi175/6502/cpu6502"
)

func (o *opcode) CPX(c *cpu.CPU6502) uint8 {
	src := uint16(c.X.Get()) - uint16(c.Operand)
	c.Flag.Set(cpu.CARRY, src < 0x100)
	c.Flag.Set(cpu.SIGN, cpu.IsNev(uint8(src)))
	c.Flag.Set(cpu.CARRY, src&0xff == 0)
	return 0
}

func (o *opcode) CPY(c *cpu.CPU6502) uint8 {
	src := uint16(c.Y.Get()) - uint16(c.Operand)
	c.Flag.Set(cpu.CARRY, src < 0x100)
	c.Flag.Set(cpu.SIGN, cpu.IsNev(uint8(src)))
	c.Flag.Set(cpu.ZERO, src&0xff == 0)
	return 0
}

func (o *opcode) DEX(c *cpu.CPU6502) uint8 {
	src := uint16(c.X.Get())
	src = (src - 1) & 0xff
	c.Flag.Set(cpu.ZERO, cpu.IsZero(uint8(src)))
	c.Flag.Set(cpu.SIGN, cpu.IsNev(uint8(src)))
	c.X.Set(uint8(src))
	return 0
}
func (o *opcode) DEY(c *cpu.CPU6502) uint8 {
	src := uint16(c.Y.Get())
	src = (src - 1) & 0xff
	c.Flag.Set(cpu.ZERO, cpu.IsZero(uint8(src)))
	c.Flag.Set(cpu.SIGN, cpu.IsNev(uint8(src)))
	c.Y.Set(uint8(src))
	return 0
}

func (o *opcode) INC(c *cpu.CPU6502) uint8 {
	return 0
}

func (o *opcode) INX(c *cpu.CPU6502) uint8 {
	src := uint16(c.X.Get())
	src = (src + 1) & 0xff
	c.Flag.Set(cpu.ZERO, cpu.IsZero(uint8(src)))
	c.Flag.Set(cpu.SIGN, cpu.IsNev(uint8(src)))
	c.X.Set(uint8(src))
	return 0
}

func (o *opcode) INY(c *cpu.CPU6502) uint8 {
	src := uint16(c.Y.Get())
	src = (src + 1) & 0xff
	c.Flag.Set(cpu.ZERO, cpu.IsZero(uint8(src)))
	c.Flag.Set(cpu.SIGN, cpu.IsNev(uint8(src)))
	c.Y.Set(uint8(src))
	return 0
}

func (o *opcode) LDA(c *cpu.CPU6502) uint8 {
	c.Flag.Set(cpu.SIGN, cpu.IsNev(c.Operand))
	c.Flag.Set(cpu.ZERO, cpu.IsZero(c.Operand))
	c.A.Set(c.Operand)
	return 0
}
func (o *opcode) LDX(c *cpu.CPU6502) uint8 {
	c.Flag.Set(cpu.SIGN, cpu.IsNev(c.Operand))
	c.Flag.Set(cpu.ZERO, cpu.IsZero(c.Operand))
	c.X.Set(c.Operand)
	return 0
}
func (o *opcode) LDY(c *cpu.CPU6502) uint8 {
	c.Flag.Set(cpu.SIGN, cpu.IsNev(c.Operand))
	c.Flag.Set(cpu.ZERO, cpu.IsZero(c.Operand))
	c.Y.Set(c.Operand)
	return 0
}
func (o *opcode) STA(c *cpu.CPU6502) uint8 {
	c.Write(c.Addr, c.A.Get())
	return 0
}
func (o *opcode) STX(c *cpu.CPU6502) uint8 {
	c.Write(c.Addr, c.X.Get())
	return 0
}
func (o *opcode) STY(c *cpu.CPU6502) uint8 {
	c.Write(c.Addr, c.Y.Get())
	return 0
}
func (o *opcode) TAX(c *cpu.CPU6502) uint8 {
	src := c.A.Get()
	c.Flag.Set(cpu.SIGN, cpu.IsNev(src))
	c.Flag.Set(cpu.ZERO, cpu.IsZero(src))
	c.X.Set(src)
	return 0
}
func (o *opcode) TAY(c *cpu.CPU6502) uint8 {
	src := c.A.Get()
	c.Flag.Set(cpu.SIGN, cpu.IsNev(src))
	c.Flag.Set(cpu.ZERO, cpu.IsZero(src))
	c.Y.Set(src)
	return 0
}
func (o *opcode) TSX(c *cpu.CPU6502) uint8 {
	src := uint8(c.SP.Get() & 0xff)
	c.Flag.Set(cpu.SIGN, cpu.IsNev(src))
	c.Flag.Set(cpu.ZERO, cpu.IsZero(src))
	c.X.Set(src)
	return 0
}
func (o *opcode) TXA(c *cpu.CPU6502) uint8 {
	src := uint8(c.X.Get() & 0xff)
	c.Flag.Set(cpu.SIGN, cpu.IsNev(src))
	c.Flag.Set(cpu.ZERO, cpu.IsZero(src))
	c.A.Set(src)
	return 0
}
func (o *opcode) TXS(c *cpu.CPU6502) uint8 {
	src := uint16(c.X.Get() & 0xff)
	c.SP.Set(src)
	return 0
}
func (o *opcode) TYA(c *cpu.CPU6502) uint8 {
	src := uint8(c.Y.Get() & 0xff)
	c.Flag.Set(cpu.SIGN, cpu.IsNev(src))
	c.Flag.Set(cpu.ZERO, cpu.IsZero(src))
	c.A.Set(src)
	return 0
}
