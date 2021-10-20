package opcode

import cpu "github.com/kabi175/6502/cpu6502"

// Clear CARRY FLAG
// Flag - C
func (*opcode) CLC(c *cpu.CPU6502) uint8 {
	c.SET_CARRY(0)
	return 0
}

// Clear DECIMAL FLAG
// Flag - D
func (*opcode) CLD(c *cpu.CPU6502) uint8 {
	c.SET_DECIMAL(0)
	return 0
}

// Clear INTERRUPT FLAG
// Flag - I
func (*opcode) CLI(c *cpu.CPU6502) uint8 {
	c.SET_INTERRUPT(0)
	return 0
}

// Clear OVERFLOW FLAG
// Flag - I
func (*opcode) CLV(c *cpu.CPU6502) uint8 {
	c.SET_OVERFLOW(false)
	return 0
}

// Set CARRY FLAG
// Flag - I
func (*opcode) SEC(c *cpu.CPU6502) uint8 {
	c.SET_CARRY(1)
	return 0
}

// Set DECIMAL FLAG
// Flag - D
func (*opcode) SED(c *cpu.CPU6502) uint8 {
	c.SET_DECIMAL(1)
	return 0
}

// Set INTERRUPT FLAG
// Flag - I
func (*opcode) SEI(c *cpu.CPU6502) uint8 {
	c.SET_INTERRUPT(1)
	return 0
}
