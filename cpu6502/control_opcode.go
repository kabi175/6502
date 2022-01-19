package cpu6502

func (o *opcode) BCC(c *CPU6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if !c.Flag.Get(CARRY) {
		c.PC.Set(c.PC.Get() + uint16(c.Operand))
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BCS(c *CPU6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if c.Flag.Get(CARRY) {
		c.PC.Set(c.PC.Get() + uint16(c.Operand))
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BEQ(c *CPU6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if c.Flag.Get(ZERO) {
		c.PC.Set(c.PC.Get() + uint16(c.Operand))
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BMI(c *CPU6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if c.Flag.Get(SIGN) {
		c.PC.Set(c.PC.Get() + uint16(c.Operand))
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BNE(c *CPU6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if !c.Flag.Get(ZERO) {
		c.PC.Set(c.PC.Get() + uint16(c.Operand))
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (o *opcode) BPL(c *CPU6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if !c.Flag.Get(SIGN) {
		c.PC.Set(c.PC.Get() + uint16(c.Operand))
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (*opcode) BVC(c *CPU6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if !c.Flag.Get(OVERFLOW) {
		c.PC.Set(c.PC.Get() + uint16(c.Operand))
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (*opcode) BVS(c *CPU6502) uint8 {
	isPageCrossed := ((c.PC.Get() & 0xFF00) != (c.PC.Get()+c.Addr)&0xFF00)
	if c.Flag.Get(OVERFLOW) {
		c.PC.Set(c.PC.Get() + uint16(c.Operand))
	}
	if isPageCrossed {
		return 2
	}
	return 1
}

func (*opcode) JMP(c *CPU6502) uint8 {
	c.PC.Set(c.Addr)
	return 0
}

func (*opcode) JSR(c *CPU6502) uint8 {
	c.PC.Decrement()
	c.PUSH(uint8(c.PC.Get() >> 8))
	c.PUSH(uint8(c.PC.Get()))
	c.PC.Set(c.Addr)
	return 0
}

func (*opcode) RTS(c *CPU6502) uint8 {
	src := uint16(c.PULL())
	src += ((uint16(c.PULL()) << 8) + 1)
	c.PC.Set(src)
	return 0
}
