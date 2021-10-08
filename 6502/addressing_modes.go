package cpu6502

/*
Immediate Addressing Mode

This a 2-byte instruction.This addressing mode have their operand defined
as next byte to the opcode.
*/
func IMM(c *cpu6502) uint8 {
	c.operand = c.fetch()
	return 0
}

/*
 */
