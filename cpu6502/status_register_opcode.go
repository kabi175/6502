package cpu6502

// Clear CARRY FLAG
// Flag - C
func (*opcode) CLC(c *CPU6502) uint8 {
	c.SET_CARRY(0)
	return 0
}

// Clear DECIMAL FLAG
// Flag - D
func (*opcode) CLD(c *CPU6502) uint8 {
	c.SET_DECIMAL(0)
	return 0
}

// Clear INTERRUPT FLAG
// Flag - I
func (*opcode) CLI(c *CPU6502) uint8 {
	c.SET_INTERRUPT(0)
	return 0
}

// Clear OVERFLOW FLAG
// Flag - I
func (*opcode) CLV(c *CPU6502) uint8 {
	c.SET_OVERFLOW(false)
	return 0
}

// Set CARRY FLAG
// Flag - I
func (*opcode) SEC(c *CPU6502) uint8 {
	c.SET_CARRY(1)
	return 0
}

// Set DECIMAL FLAG
// Flag - D
func (*opcode) SED(c *CPU6502) uint8 {
	c.SET_DECIMAL(1)
	return 0
}

// Set INTERRUPT FLAG
// Flag - I
func (*opcode) SEI(c *CPU6502) uint8 {
	c.SET_INTERRUPT(1)
	return 0
}
