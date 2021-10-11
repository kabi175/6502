package opcode

import cpu "github.com/kabi175/6502/cpu6502"

// Force Break
// Flags - I, B
func BRK(c *cpu.Cpu6502) uint8 {
	c.PC.Increment()
	c.PUSH(uint8(c.PC.Get() >> 8))
	c.PUSH(uint8(c.PC.Get() & 0xff))
	c.SET_BREAk(1)
	c.PUSH(c.Flag.Byte())
	c.SET_INTERRUPT(1)
	c.PC.Set(uint16(c.Read(0xFFFE)) | uint16(c.Read(0xFFFF))<<8)
	return 0
}

// Return from Interrupt
// Flags - nil
func RTI(c *cpu.Cpu6502) uint8 {
	src := uint16(c.PULL())
	c.SP.Set(src)
	src = uint16(c.PULL())
	src |= uint16(c.PULL()) << 8
	return 0
}

// NO Operations
func NOP(c *cpu.Cpu6502) uint8 {
	return 0
}
