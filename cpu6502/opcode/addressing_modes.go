package opcode

import cpu "github.com/kabi175/6502/cpu6502"

/* Accumulator Mode*
In this mode, instruction operaties on data on accumulator.
No operations are needed.
*/
func ACC(c *cpu.Cpu6502) uint8 {
	// No operations are  needed
	return 0
}

/* Immediate Addressing Mode

This a 2-byte instruction.This addressing mode have their operand defined
as next byte to the opcode.
*/
func IMM(c *cpu.Cpu6502) uint8 {
	c.Operand = c.Fetch()
	return 0
}

/* Zero-Page or Zero-Page Absolute Addressing Mode

This  a 2-byte instruction.
This mode specifies the location of the operand.
This mode is capable of addressing only 1st 256 bytes of memery.

1st byte -> Instruction

2nd byte -> addr low

By default high addr will 0x00 -> Zero page
*/
func ZPA(c *cpu.Cpu6502) uint8 {
	high := uint8(0x00)
	low := c.Fetch()
	addr := cpu.LittleEndianAddr(low, high)
	c.Addr = addr
	c.Operand = c.Read(addr)
	return 0
}

/* Zero-Page X
This is 2 byte opcode.
This mode specifies the location of the operand.
This mode is capable of addressing only 1st 256 bytes of memory.

1st byte -> instruction

2nd byte -> addr low
addr = low + X
*/
func ZPX(c *cpu.Cpu6502) uint8 {
	low := c.Fetch()
	addr := (uint16(low) + uint16(c.X.Get())) & 0x00ff
	c.Addr = addr
	c.Operand = c.Read(addr)
	return 0
}

/* Zero-Page Y Mode

This is 2 byte opcode.
This mode specifies the location of the operand.
This mode is capable of addressing only 1st 256 bytes of memory.

1st byte -> instruction

2nd byte -> addr low

addr = low + Y
*/
func ZPY(c *cpu.Cpu6502) uint8 {
	low := c.Fetch()
	addr := (uint16(low) + uint16(c.Y.Get())) & 0x00ff
	c.Addr = addr
	c.Operand = c.Read(addr)
	return 0
}

/* Absolute Addressing Mode

This  a 3-byte instruction. This mode specifies the location of the operand.

1st byte -> Instruction

2nd byte -> addr low

3rd byte -> addr high

NOTE: 6502 uses Little Endian
*/
func ABS(c *cpu.Cpu6502) uint8 {
	low := c.Fetch()
	high := c.Fetch()
	addr := cpu.LittleEndianAddr(low, high)
	c.Addr = addr
	c.Operand = c.Read(addr)
	return 0
}

/* Absolute X mode

This mode contains is a 3-byte opcode. This mode specifies the location of the operand.

1st byte -> Instruction

2nd byte -> low byte

3rd byte -> high byte

Add the X register with the address produced by (high,low). If the operations crossed the  page then it consumes one more additional cycle.

*/
func ABX(c *cpu.Cpu6502) uint8 {
	low := c.Fetch()
	high := c.Fetch()
	addr := cpu.LittleEndianAddr(low, high) + uint16(c.X.Get())
	c.Addr = addr
	c.Operand = c.Read(addr)

	// Check for page cross
	if uint16(c.X.Get())+uint16(low) > 0x00ff {
		return 1
	}

	return 0
}

/* Absolute Y mode

This mode contains is a 3-byte opcode. This mode specifies the location of the operand.

1st byte -> Instruction

2nd byte -> low byte

3rd byte -> high byte

Add the Y register with the address produced by (high,low). If the operations crossed the  page then it consumes one more additional cycle.
*/
func ABY(c *cpu.Cpu6502) uint8 {
	low := c.Fetch()
	high := c.Fetch()
	addr := cpu.LittleEndianAddr(low, high) + uint16(c.Y.Get())
	c.Addr = addr
	c.Operand = c.Read(addr)

	// Check for page cross
	if uint16(c.Y.Get())+uint16(low) > 0x00ff {
		return 1
	}

	return 0
}

/* Implied Addressing Mode

No operand addresses are required for this mode. They are implied by the
instruction.
*/
func IMP(c *cpu.Cpu6502) uint8 {
	// No logic is required
	return 0
}

// Relative Mode
func REL(c *cpu.Cpu6502) uint8 {
	c.Operand = c.Fetch()
	return 0
}

// Indirect X Or Indexed Indirect
//
/*
This a 2 byte opcode.
1st byte contains the instruction. 2nd byte contains the addres in the zero page.
Add the Zero-page address with X register to get a new address in the Zero-page.
This new address contains the target address.
*/
func IDX(c *cpu.Cpu6502) uint8 {
	zpAddr := uint16(c.Fetch()) + uint16(c.X.Get())
	zpAddr &= 0x00ff
	low := c.Read(zpAddr)
	high := c.Read(zpAddr + 1)
	targetAddr := cpu.LittleEndianAddr(low, high)
	c.Addr = targetAddr
	c.Operand = c.Read(targetAddr)
	return 0
}

// Indirect Y or Indirect Indexed
//
// This is a 2 byte opcode. 1st byte is the instruction.
// 2nd byte is the Zero-page memory location. Featch the address stored at Zero-page address
// and Added it with Y Resister to get final target address.
func IDY(c *cpu.Cpu6502) uint8 {
	zpAddr := uint16(c.Fetch())
	low := c.Read(zpAddr)
	high := c.Read(zpAddr + 1)

	targetAddr := cpu.LittleEndianAddr(low, high) + uint16(c.Y.Get())

	c.Addr = targetAddr
	c.Operand = c.Read(targetAddr)

	return 0
}

// Indirect Addressing
//
// The JMP instruction is the only instruction that uses this addressing mode.
// It is a 3 byte instruction - the 2nd and 3rd bytes are an absolute address.
// The set the PC to the address stored at that address.
func IND(c *cpu.Cpu6502) uint8 {
	low := uint16(c.Fetch())
	high := uint16(c.Fetch()) << 8
	addr := low | high
	targetLow := c.Read(addr)
	targetHigh := c.Read(addr + 1)
	c.Addr = cpu.LittleEndianAddr(targetLow, targetHigh)
	return 0
}