package cpu6502

// This instruction adds memory and carry to the A Register.
// M + C + A -> A
// if DECIMAL Flag enable perform BCD addition
// Flags: C, V, S, Z
func (o *opcode) ADC(c *CPU6502) uint8 {
	carry := uint8(0)
	if c.Flag.Get(CARRY) {
		carry = 1
	}
	temp := uint16(c.A.Get()) + uint16(c.operand) + uint16(carry)
	if c.Flag.Get(DECIMAL) {
		if (c.A.Get()&0x0f)+(c.operand&0xf)+carry > 0x09 {
			temp += 6
		}
		c.SET_OVERFLOW(((c.A.Get()^c.operand)&0x80 == 0) && ((uint16(c.A.Get())^temp)&0x80) != 0)
		if temp > 0x99 {
			temp += 96
		}
		c.SET_CARRY(0)
		if temp > 0x99 {
			c.SET_CARRY(1)
		}
	} else {
		c.SET_OVERFLOW(((c.A.Get()^c.operand)&0x80 == 0) && ((uint16(c.A.Get())^temp)&0x80) != 0)
		c.SET_CARRY(temp & 0xff00)
	}
	c.SET_SIGN(temp)
	c.SET_ZERO(temp & 0xff)
	c.A.Set(uint8(temp))
	return 0
}

// This is symbolically represented by A & M -> A.
// Flags: Z, S
func (o *opcode) AND(c *CPU6502) uint8 {
	src := c.operand & c.A.Get()
	c.SET_SIGN(uint16(src))
	c.SET_ZERO(uint16(src))
	c.A.Set(src)
	return 0
}

// Shift 1 bit to left (memory or Reg)
// STORE opearnd in memory or accumulator depending on addressing mode.
// Flags: C, Z, S
func (o *opcode) ASL(c *CPU6502) uint8 {
	c.SET_CARRY(uint16(c.operand & 0x80))
	c.operand <<= 1
	c.SET_SIGN(uint16(c.operand))
	c.SET_ZERO(uint16(c.operand))
	// Store Operand in memory or in A reg
	if o.Mode == ACC {
		c.A.Set(c.operand)
		return 0
	}
	c.Write(c.addr, c.operand)
	return 0
}

// Test Bits in Memory with Accumulator
// FLags: O, Z, S
func (o *opcode) BIT(c *CPU6502) uint8 {
	c.SET_SIGN(uint16(c.operand))
	c.SET_ZERO(uint16(c.operand) & uint16(c.A.Get()))
	c.SET_OVERFLOW(c.operand&0x40 != 0)
	return 0
}

// Compare Memory and Accumulator
// Flags: S, Z, C
func (o *opcode) CMP(c *CPU6502) uint8 {
	src := uint16(c.A.Get()) - uint16(c.operand)
	c.SET_CARRY(0)
	if src < 0x100 {
		c.SET_CARRY(1)
	}
	c.SET_SIGN(src)
	c.SET_ZERO(src & 0x00ff)
	return 0
}

// Decrement Memory by One
// Flags S, C
func (o *opcode) DEC(c *CPU6502) uint8 {
	src := (c.operand - 1) & 0xff
	c.SET_SIGN(uint16(src))
	c.SET_ZERO(uint16(src))
	c.Write(c.addr, src)
	return 0
}

// XOR Memory with Accumulator
// Flags S, C
func (o *opcode) EOR(c *CPU6502) uint8 {
	src := (c.operand ^ c.A.Get())
	c.SET_SIGN(uint16(src))
	c.SET_ZERO(uint16(src))
	c.A.Set(src)
	return 0
}

// Shift Right One Bit (Memory or Accumulator)
//Flags C, Z, S
func (o *opcode) LSR(c *CPU6502) uint8 {
	c.SET_CARRY(uint16(c.operand & 0x01))
	c.operand >>= 1
	c.SET_SIGN(uint16(c.operand))
	c.SET_ZERO(uint16(c.operand))
	// Store Operand in memory or in A reg
	if o.Mode == ACC {
		c.A.Set(c.operand)
		return 0
	}
	c.Write(c.addr, c.operand)
	return 0
}

// OR memory with A reg
// Flags S, Z
func (o *opcode) ORA(c *CPU6502) uint8 {
	c.operand |= c.A.Get()
	c.SET_SIGN(uint16(c.operand))
	c.SET_ZERO(uint16(c.operand))
	c.A.Set(c.operand)
	return 0
}

// Rotate One Bit Left (Memory or Accumulator)
// Flags S, Z, C
func (o *opcode) ROL(c *CPU6502) uint8 {
	src := uint16(c.operand) << 1
	if c.Flag.Get(CARRY) {
		src |= 0x1
	}
	c.SET_CARRY(0)
	if src > 0xFF {
		c.SET_CARRY(1)
	}
	src &= 0xff
	c.SET_SIGN(src)
	c.SET_ZERO(src)
	// Store Operand in memory or in A reg
	if o.Mode == ACC {
		c.A.Set(uint8(src))
		return 0
	}
	c.Write(c.addr, uint8(src))
	return 0
}

// Rotate One Bit Right (Memory or Accumulator)
// Flags S, Z, C
func (o *opcode) ROR(c *CPU6502) uint8 {
	src := uint16(c.operand)
	if c.Flag.Get(CARRY) {
		src |= 0x100
	}
	c.SET_CARRY(src & 0x01)
	src >>= 1
	c.SET_SIGN(src)
	c.SET_ZERO(src)
	// Store Operand in memory or in A reg
	if o.Mode == ACC {
		c.A.Set(uint8(src))
		return 0
	}
	c.Write(c.addr, uint8(src))
	return 0
}

// Subtract Memory from Accumulator with Borrow
// Flags S, Z, V, C
func (o *opcode) SBC(c *CPU6502) uint8 {
	carry := 1
	if c.Flag.Get(CARRY) {
		carry = 0
	}
	temp := uint16(c.A.Get()) - uint16(c.operand) - uint16(carry)
	c.SET_SIGN(temp)
	c.SET_ZERO(temp & 0xff)
	c.SET_OVERFLOW((((uint16(c.A.Get())^temp)&0x80) != 0 && ((c.A.Get()^c.operand)&0x80) != 0))
	if c.Flag.Get(DECIMAL) {
		carry := uint8(0)
		if !c.Flag.Get(CARRY) {
			carry++
		}
		if ((c.A.Get() & 0xf) - carry) < (c.operand & 0xf) {
			temp -= 6
		}
		if temp > 0x99 {
			temp -= 0x60
		}
	}
	c.SET_CARRY(0)
	if temp < 0x100 {
		c.SET_CARRY(1)
	}
	c.A.Set(uint8(temp & 0xff))
	return 0
}
