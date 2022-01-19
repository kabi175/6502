package cpu6502

func (o *opcode) CPX(c *CPU6502) uint8 {
	src := uint16(c.X.Get()) - uint16(c.Operand)
	c.Flag.Set(CARRY, src < 0x100)
	c.Flag.Set(SIGN, IsNev(uint8(src)))
	c.Flag.Set(CARRY, src&0xff == 0)
	return 0
}

func (o *opcode) CPY(c *CPU6502) uint8 {
	src := uint16(c.Y.Get()) - uint16(c.Operand)
	c.Flag.Set(CARRY, src < 0x100)
	c.Flag.Set(SIGN, IsNev(uint8(src)))
	c.Flag.Set(ZERO, src&0xff == 0)
	return 0
}

func (o *opcode) DEX(c *CPU6502) uint8 {
	src := uint16(c.X.Get())
	src = (src - 1) & 0xff
	c.Flag.Set(ZERO, IsZero(uint8(src)))
	c.Flag.Set(SIGN, IsNev(uint8(src)))
	c.X.Set(uint8(src))
	return 0
}
func (o *opcode) DEY(c *CPU6502) uint8 {
	src := uint16(c.Y.Get())
	src = (src - 1) & 0xff
	c.Flag.Set(ZERO, IsZero(uint8(src)))
	c.Flag.Set(SIGN, IsNev(uint8(src)))
	c.Y.Set(uint8(src))
	return 0
}

func (o *opcode) INC(c *CPU6502) uint8 {
	return 0
}

func (o *opcode) INX(c *CPU6502) uint8 {
	src := uint16(c.X.Get())
	src = (src + 1) & 0xff
	c.Flag.Set(ZERO, IsZero(uint8(src)))
	c.Flag.Set(SIGN, IsNev(uint8(src)))
	c.X.Set(uint8(src))
	return 0
}

func (o *opcode) INY(c *CPU6502) uint8 {
	src := uint16(c.Y.Get())
	src = (src + 1) & 0xff
	c.Flag.Set(ZERO, IsZero(uint8(src)))
	c.Flag.Set(SIGN, IsNev(uint8(src)))
	c.Y.Set(uint8(src))
	return 0
}

func (o *opcode) LDA(c *CPU6502) uint8 {
	c.Flag.Set(SIGN, IsNev(c.Operand))
	c.Flag.Set(ZERO, IsZero(c.Operand))
	c.A.Set(c.Operand)
	return 0
}
func (o *opcode) LDX(c *CPU6502) uint8 {
	c.Flag.Set(SIGN, IsNev(c.Operand))
	c.Flag.Set(ZERO, IsZero(c.Operand))
	c.X.Set(c.Operand)
	return 0
}
func (o *opcode) LDY(c *CPU6502) uint8 {
	c.Flag.Set(SIGN, IsNev(c.Operand))
	c.Flag.Set(ZERO, IsZero(c.Operand))
	c.Y.Set(c.Operand)
	return 0
}
func (o *opcode) STA(c *CPU6502) uint8 {
	c.Write(c.Addr, c.A.Get())
	return 0
}
func (o *opcode) STX(c *CPU6502) uint8 {
	c.Write(c.Addr, c.X.Get())
	return 0
}
func (o *opcode) STY(c *CPU6502) uint8 {
	c.Write(c.Addr, c.Y.Get())
	return 0
}
func (o *opcode) TAX(c *CPU6502) uint8 {
	src := c.A.Get()
	c.Flag.Set(SIGN, IsNev(src))
	c.Flag.Set(ZERO, IsZero(src))
	c.X.Set(src)
	return 0
}
func (o *opcode) TAY(c *CPU6502) uint8 {
	src := c.A.Get()
	c.Flag.Set(SIGN, IsNev(src))
	c.Flag.Set(ZERO, IsZero(src))
	c.Y.Set(src)
	return 0
}
func (o *opcode) TSX(c *CPU6502) uint8 {
	src := uint8(c.SP.Get() & 0xff)
	c.Flag.Set(SIGN, IsNev(src))
	c.Flag.Set(ZERO, IsZero(src))
	c.X.Set(src)
	return 0
}
func (o *opcode) TXA(c *CPU6502) uint8 {
	src := uint8(c.X.Get() & 0xff)
	c.Flag.Set(SIGN, IsNev(src))
	c.Flag.Set(ZERO, IsZero(src))
	c.A.Set(src)
	return 0
}
func (o *opcode) TXS(c *CPU6502) uint8 {
	src := uint16(c.X.Get() & 0xff)
	c.SP.Set(src)
	return 0
}
func (o *opcode) TYA(c *CPU6502) uint8 {
	src := uint8(c.Y.Get() & 0xff)
	c.Flag.Set(SIGN, IsNev(src))
	c.Flag.Set(ZERO, IsZero(src))
	c.A.Set(src)
	return 0
}
