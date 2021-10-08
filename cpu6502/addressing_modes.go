package cpu6502

/* Accumulator Mode*
In this mode, instruction operaties on data on accumulator.
No operations are needed.
*/
func ACC(c *cpu6502) uint8 {
	// No operations are  needed
	return 0
}

/* Immediate Addressing Mode

This a 2-byte instruction.This addressing mode have their operand defined
as next byte to the opcode.
*/
func IMM(c *cpu6502) uint8 {
	c.operand = c.fetch()
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
func ZPA(c *cpu6502) uint8 {
	high := uint8(0x00)
	low := c.fetch()
	addr := littleEndianAddr(low, high)
	c.addr = addr
	c.operand = c.read(addr)
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
func ZPX(c *cpu6502) uint8 {
	low := c.fetch()
	addr := (uint16(low) + uint16(c.X.Get())) & 0x00ff
	c.addr = addr
	c.operand = c.read(addr)
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
func ZPY(c *cpu6502) uint8 {
	low := c.fetch()
	addr := (uint16(low) + uint16(c.Y.Get())) & 0x00ff
	c.addr = addr
	c.operand = c.read(addr)
	return 0
}

/* Absolute Addressing Mode

This  a 3-byte instruction. This mode specifies the location of the operand.

1st byte -> Instruction

2nd byte -> addr low

3rd byte -> addr high

NOTE: 6502 uses Little Endian
*/
func ABS(c *cpu6502) uint8 {
	low := c.fetch()
	high := c.fetch()
	addr := littleEndianAddr(low, high)
	c.addr = addr
	c.operand = c.read(addr)
	return 0
}

/* Absolute X mode

This mode contains is a 3-byte opcode. This mode specifies the location of the operand.

1st byte -> Instruction

2nd byte -> low byte

3rd byte -> high byte

Add the X register with the address produced by (high,low). If the operations crossed the  page then it consumes one more additional cycle.

*/
func ABX(c *cpu6502) uint8 {
	low := c.fetch()
	high := c.fetch()
	addr := littleEndianAddr(low, high) + uint16(c.X.Get())
	c.addr = addr
	c.operand = c.read(addr)

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
func ABY(c *cpu6502) uint8 {
	low := c.fetch()
	high := c.fetch()
	addr := littleEndianAddr(low, high) + uint16(c.Y.Get())
	c.addr = addr
	c.operand = c.read(addr)

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
func IMP(c *cpu6502) uint8 {
	// No logic is required
	return 0
}

// Relative Mode
func REL(c *cpu6502) uint8 {
	c.operand = c.fetch()
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
func IDX(c *cpu6502) uint8 {
	zpAddr := uint16(c.fetch()) + uint16(c.X.Get())
	zpAddr &= 0x00ff
	low := c.read(zpAddr)
	high := c.read(zpAddr + 1)
	targetAddr := littleEndianAddr(low, high)
	c.addr = targetAddr
	c.operand = c.read(targetAddr)
	return 0
}

// Indirect Y or Indirect Indexed
//
// This is a 2 byte opcode. 1st byte is the instruction.
// 2nd byte is the Zero-page memory location. Featch the address stored at Zero-page address
// and Added it with Y Resister to get final target address.
func IDY(c *cpu6502) uint8 {
	zpAddr := uint16(c.fetch())
	low := c.read(zpAddr)
	high := c.read(zpAddr + 1)

	targetAddr := littleEndianAddr(low, high) + uint16(c.Y.Get())

	c.addr = targetAddr
	c.operand = c.read(targetAddr)

	return 0
}

// Indirect Addressing
//
// The JMP instruction is the only instruction that uses this addressing mode.
// It is a 3 byte instruction - the 2nd and 3rd bytes are an absolute address.
// The set the PC to the address stored at that address.
func IND(c *cpu6502) uint8 {
	low := uint16(c.fetch())
	high := uint16(c.fetch()) << 8
	addr := low | high
	targetLow := c.read(addr)
	targetHigh := c.read(addr + 1)
	c.addr = littleEndianAddr(targetLow, targetHigh)
	return 0
}
