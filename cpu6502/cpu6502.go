package cpu6502

import (
	"github.com/kabi175/6502/model"
)

type CPU6502 struct {
	Operand       uint8
	Addr          uint16
	opcodeBuilder model.OpcodeBuilder
	PC            model.PC16
	SP            model.SP8
	Flag          model.FlagRegister
	A             model.GP8
	X             model.GP8
	Y             model.GP8
	Bus           model.Bus16
	eventQueue    model.EventQueue
	deb           model.Debugger
}

func (c *CPU6502) AttachOpcodeBuilder(builder model.OpcodeBuilder) *CPU6502 {
	c.opcodeBuilder = builder
	return c
}

// Helper function to check the channel for new message or chan close
func isClosed(close chan bool) bool {
	select {
	case <-close:
		return true
	default:
		return false
	}
}

/*
 -> Check for cpu quit
 -> Fetch Opcode from memory
 -> Get instructon set with opcode
 -> Increment Program Counter
 -> Execute AddressMode func to get operand
 -> Execute the instruction func with operand
 -> Track no.of cycles that took for execution
*/
func (c *CPU6502) Execute(close chan bool) {
	c.Reset()
	for !isClosed(close) {
		opcodeHex := c.Fetch()
		opcode := c.opcodeBuilder.Build(opcodeHex)
		opcode.Execute()
		if opcode.IsBreak() || (c != nil && c.deb.IsEnd(c.PC.Get())) {
			break
		}
	}
}

func (c *CPU6502) Read(add uint16) uint8 {
	return c.Bus.Read(add)
}

func (c *CPU6502) Write(add uint16, data uint8) {
	c.Bus.Write(add, data)
}

func (c *CPU6502) Fetch() uint8 {
	data := c.Read(c.PC.Get())
	c.PC.Increment()
	return data
}

func (c *CPU6502) Reset() {
	// Reset GP registers
	c.A.Set(0x00)
	c.X.Set(0x00)
	c.Y.Set(0x00)

	// Reset Flags
	c.Flag.Reset()

	// Reset Program Counter to the
	low := c.Read(0xFFFC)
	high := c.Read(0xFFFC + 1)
	addr := LittleEndianAddr(low, high)
	c.PC.Set(addr)
}

// Set Overflow Flag on condition
func (c *CPU6502) SET_OVERFLOW(cond bool) {
	c.Flag.Set(OVERFLOW, cond)
}

// Set CARRY if val is non zero
func (c *CPU6502) SET_CARRY(val uint16) {
	c.Flag.Set(CARRY, val != 0)
}

// Set SIGN if val bit 7 is set
func (c *CPU6502) SET_SIGN(val uint16) {
	c.Flag.Set(SIGN, val&0x80 != 0)
}

// Set Zero if val is zero
func (c *CPU6502) SET_ZERO(val uint16) {
	c.Flag.Set(ZERO, val == 0)
}

// Set BREAk if val is non zero
func (c *CPU6502) SET_BREAk(val uint16) {
	c.Flag.Set(BREAK, val != 0)
}

// Set INTERRUPT  if val is non zero
func (c *CPU6502) SET_INTERRUPT(val uint16) {
	c.Flag.Set(INTERRUPT, val != 0)
}

// Set DECIMAL if val is non zero
func (c *CPU6502) SET_DECIMAL(val uint16) {
	c.Flag.Set(DECIMAL, val != 0)
}

func (c *CPU6502) PUSH(byte uint8) {
	c.Write(0x0100|c.SP.Get(), byte)
	c.SP.Decrement()
}

func (c *CPU6502) PULL() uint8 {
	c.SP.Increment()
	return c.Read(0x0100 | c.SP.Get())
}
