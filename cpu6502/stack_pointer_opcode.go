package cpu6502

// Push Accumulator on Stack
// Flags - Nil
func (*opcode) PHA(c *CPU6502) uint8 {
	c.PUSH(c.A.Get())
	return 0
}

// Push Processor Status on Stack
// Flags - Nil
func (*opcode) PHP(c *CPU6502) uint8 {
	src := c.Flag.Byte()
	c.PUSH(src)
	return 0
}

// Pull Accumulator from Stack
// Flags - S, Z
func (*opcode) PLA(c *CPU6502) uint8 {
	src := c.PULL()
	c.SET_SIGN(uint16(src))
	c.SET_ZERO(uint16(src))
	return 0
}

// Pull Processor Status from Stack
func (*opcode) PLP(c *CPU6502) uint8 {
	src := c.PULL()
	c.Flag.SetByte(src)
	return 0
}
