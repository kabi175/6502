package opcode

import cpu "github.com/kabi175/6502/cpu6502"

func (o *opcode) BCC(c *cpu.Cpu6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if !c.Flag.Get(cpu.CARRY) {
		c.PC.Set(c.PC.Get() + c.Addr)
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BCS(c *cpu.Cpu6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if c.Flag.Get(cpu.CARRY) {
		c.PC.Set(c.PC.Get() + c.Addr)
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BEQ(c *cpu.Cpu6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if c.Flag.Get(cpu.ZERO) {
		c.PC.Set(c.PC.Get() + c.Addr)
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BMI(c *cpu.Cpu6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if c.Flag.Get(cpu.SIGN) {
		c.PC.Set(c.PC.Get() + c.Addr)
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BNE(c *cpu.Cpu6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if !c.Flag.Get(cpu.ZERO) {
		c.PC.Set(c.PC.Get() + c.Addr)
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BPL(c *cpu.Cpu6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if !c.Flag.Get(cpu.SIGN) {
		c.PC.Set(c.PC.Get() + c.Addr)
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (*opcode) BVC(c *cpu.Cpu6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if !c.Flag.Get(cpu.OVERFLOW) {
		c.PC.Set(c.PC.Get() + c.Addr)
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (*opcode) BVS(c *cpu.Cpu6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if c.Flag.Get(cpu.OVERFLOW) {
		c.PC.Set(c.PC.Get() + c.Addr)
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (*opcode) JMP(c *cpu.Cpu6502) uint8 {
	c.PC.Set(c.Addr)
	return 0
}

func (*opcode) JSR(c *cpu.Cpu6502) uint8 {
	c.PC.Decrement()
	c.Push(uint8(c.PC.Get() >> 8))
	c.Push(uint8(c.PC.Get()))
	c.PC.Set(c.Addr)
	return 0
}

func (*opcode) RTS(c *cpu.Cpu6502) uint8 {
	src := uint16(c.Pull())
	src += ((uint16(c.Pull()) << 8) + 1)
	c.PC.Set(src)
	return 0
}
