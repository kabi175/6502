package opcode

import cpu "github.com/kabi175/6502/cpu6502"

// This instruction adds memory and carry to the A Register.
// M + C + A -> A
// if DECIMAL Flag enable perform BCD addition
// Flags: C, V, S, Z
func (o *opcode) ADC(c *cpu.Cpu6502) uint8 {
	mem := c.Operand
	A := c.A.Get()
	if !c.Flag.Get(cpu.DECIMAL) {
		// Process the addition in 8-bit binary addion mode
		if c.Flag.Get(cpu.CARRY) {
			mem += 0x01
		}
		result := A + mem

		c.A.Set(result)

		c.Flag.Set(cpu.OVERFLOW, cpu.IsOverFlow(A, mem))
		c.Flag.Set(cpu.CARRY, cpu.IsCarry(A, mem))
		c.Flag.Set(cpu.SIGN, cpu.IsNev(result))
		c.Flag.Set(cpu.ZERO, cpu.IsZero(result))

		return 0
	}
	// process addition in Decimal Mode
	// Split the memory and A reg
	memLow := mem & 0x0f
	memHigh := (mem & 0xf0) >> 4
	ALow := A & 0x0f
	AHigh := (A & 0xf0) >> 4

	// Add memory low and A Register low to get result low
	resultLow := memLow + ALow

	// if carry flag set add 1 to low result
	if c.Flag.Get(cpu.CARRY) {
		resultLow++
	}

	// Add memory high and A Register high to get result high
	resultHigh := memHigh + AHigh

	// if result  low greater than 9 (to fin in a digit),
	// add the carry to result high
	if resultLow > 9 {
		resultHigh++
	}
	// use result  low 4 bit and high 4 bit to get result
	result := resultLow & 0x0f
	result |= ((resultHigh & 0x0f) << 4)

	// update the result value to the A reg
	c.A.Set(result)

	// if result high is greate that 9, set CARRY Flag
	// if result high's 1st(MSB) 4 bit is set, an overflow is caused
	// if the result is zero set ZERO Flag
	// if the result's bit-7 is set, then set SIGN Flag
	c.Flag.Set(cpu.CARRY, resultHigh > 9)
	c.Flag.Set(cpu.OVERFLOW, resultHigh&0xf0 != 0)
	c.Flag.Set(cpu.ZERO, cpu.IsZero(result))
	c.Flag.Set(cpu.SIGN, cpu.IsNev(result))

	return 0
}

// This is symbolically represented by A & M -> A.
// Flags: Z, S
func (o *opcode) AND(c *cpu.Cpu6502) uint8 {
	src := c.Operand & c.A.Get()
	c.SET_SIGN(uint16(src))
	c.SET_ZERO(uint16(src))
	c.A.Set(src)
	return 0
}

// Shift 1 bit to left (memory or Reg)
// STORE opearnd in memory or accumulator depending on addressing mode.
// Flags: C, Z, S
func (o *opcode) ASL(c *cpu.Cpu6502) uint8 {
	c.SET_CARRY(uint16(c.Operand & 0x80))
	c.Operand <<= 1
	c.SET_SIGN(uint16(c.Operand))
	c.SET_ZERO(uint16(c.Operand))
	// Store Operand in memory or in A reg
	if o.mode == "ACC" {
		c.A.Set(c.Operand)
		return 0
	}
	c.Write(c.Addr, c.Operand)
	return 0
}

// Test Bits in Memory with Accumulator
// FLags: O, Z, S
func (o *opcode) BIT(c *cpu.Cpu6502) uint8 {
	c.SET_SIGN(uint16(c.Operand))
	c.SET_ZERO(uint16(c.Operand))
	c.SET_OVERFLOW(c.Operand&0x40 != 0)
	return 0
}

// Flags: S, Z, C
func (o *opcode) CMP(c *cpu.Cpu6502) uint8 {
	src := uint16(c.A.Get()) - uint16(c.Operand)
	c.SET_CARRY(src & 0x0100)
	c.SET_SIGN(src)
	c.SET_ZERO(src & 0x00ff)
	return 0
}

// Decrement Memory by One
func (O *opcode) DEC(c *cpu.Cpu6502) uint8 {
	src := (c.Operand - 1) & 0xff
	c.SET_SIGN(uint16(src))
	c.SET_ZERO(uint16(src))
	c.Write(c.Addr, src)
	return 0
}

// XOR Memory with Accumulator
func (o *opcode) EOR(c *cpu.Cpu6502) uint8 {
	src := (c.Operand ^ c.A.Get())
	c.SET_SIGN(uint16(src))
	c.SET_ZERO(uint16(src))
	c.Write(c.Addr, src)
	return 0
}

func (o *opcode) LSR(c *cpu.Cpu6502) uint8 {
	c.SET_CARRY(uint16(c.Operand & 0x01))
	c.Operand <<= 1
	c.SET_SIGN(uint16(c.Operand))
	c.SET_ZERO(uint16(c.Operand))
	// Store Operand in memory or in A reg
	if o.mode == "ACC" {
		c.A.Set(c.Operand)
		return 0
	}
	c.Write(c.Addr, c.Operand)
	return 0
}

// OR memory with A reg
func (o *opcode) ORA(c *cpu.Cpu6502) uint8 {
	c.Operand |= c.A.Get()
	c.SET_SIGN(uint16(c.Operand))
	c.SET_ZERO(uint16(c.Operand))
	c.A.Set(c.Operand)
	return 0
}

//
func (o *opcode) ROL(c *cpu.Cpu6502) uint8 {
	src := uint16(c.Operand) << 1
	if c.Flag.Get(cpu.CARRY) {
		src |= 0x1
	}
	c.SET_CARRY(src & 0x00)
	src &= 0xff
	c.SET_SIGN(uint16(c.Operand))
	c.SET_ZERO(uint16(c.Operand))
	// Store Operand in memory or in A reg
	if o.mode == "ACC" {
		c.A.Set(c.Operand)
		return 0
	}
	c.Write(c.Addr, c.Operand)
	return 0
}

func (o *opcode) ROR(c *cpu.Cpu6502) uint8 {
	src := uint16(c.Operand)
	if c.Flag.Get(cpu.CARRY) {
		src |= 0x100
	}
	c.SET_CARRY(src & 0x01)
	src >>= 1
	c.SET_SIGN(uint16(c.Operand))
	c.SET_ZERO(uint16(c.Operand))
	// Store Operand in memory or in A reg
	if o.mode == "ACC" {
		c.A.Set(c.Operand)
		return 0
	}
	c.Write(c.Addr, c.Operand)
	return 0
}

func (o *opcode) SBC(c *cpu.Cpu6502) uint8 {
	temp := uint16(c.A.Get()) - uint16(c.Operand)
	if !c.Flag.Get(cpu.CARRY) {
		temp++
	}
	c.SET_SIGN(temp)
	c.SET_ZERO(temp & 0xff)
	c.SET_OVERFLOW((((uint16(c.A.Get())^temp)&0x80) != 0 && ((c.A.Get()^c.Operand)&0x80) != 0))
	if c.Flag.Get(cpu.DECIMAL) {
		carry := uint8(0)
		if !c.Flag.Get(cpu.CARRY) {
			carry++
		}
		if ((c.A.Get() & 0xf) - carry) < (c.Operand & 0xf) {
			temp -= 6
		}
		if temp > 0x99 {
			temp -= 0x60
		}
	}
	c.SET_CARRY(temp & 0x00)
	c.A.Set(uint8(temp & 0xff))
	return 0
}
