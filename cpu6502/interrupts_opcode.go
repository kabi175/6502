package cpu6502

// Force Break
// Flags - I, B
func (*opcode) BRK(c *CPU6502) uint8 {
	c.PC.Increment()
	c.push(uint8(c.PC.Get() >> 8))
	c.push(uint8(c.PC.Get() & 0xff))
	c.SET_BREAk(1)
	c.push(c.Flag.Byte())
	c.SET_INTERRUPT(1)
	c.PC.Set(uint16(c.Read(0xFFFE)) | uint16(c.Read(0xFFFF))<<8)
	return 0
}

// Return from Interrupt
// Flags - nil
func (*opcode) RTI(c *CPU6502) uint8 {
	src := uint16(c.pull())
	c.SP.Set(src)
	src = uint16(c.pull())
	src |= uint16(c.pull()) << 8
	return 0
}

// NO Operations
func (*opcode) NOP(c *CPU6502) uint8 {
	return 0
}
