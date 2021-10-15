package cpu6502

type Cpu6502 struct {
	Operand       uint8
	Addr          uint16
	OpcodeBuilder func(uint8) Opcode
	PC            PC16
	SP            SP8
	Flag          FlagRegister
	A             GP8
	X             GP8
	Y             GP8
	Bus           Bus16
	deb           Debugger
}

//	Constructor function to create CPU6502
func New(bus Bus16, opcodeBuilder func(uint8) Opcode, deb Debugger) *Cpu6502 {
	return &Cpu6502{
		OpcodeBuilder: opcodeBuilder,
		PC:            NewPC(),
		SP:            NewSP8(),
		Flag:          NewFlagRegister(),
		A:             NewGP8(),
		X:             NewGP8(),
		Y:             NewGP8(),
		Bus:           bus,
		deb:           deb,
	}
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
func (c *Cpu6502) Execute(close chan bool) {
	c.Reset()
	for !isClosed(close) {
		opcodeHex := c.Fetch()
		opcode := c.OpcodeBuilder(opcodeHex)
		opcode.Execute(c)
		c.deb.Render(c)
		if opcode.IsBreak() {
			break
		}
	}
}

func (c *Cpu6502) Read(add uint16) uint8 {
	return c.Bus.Read(add)
}

func (c *Cpu6502) Write(add uint16, data uint8) {
	c.Bus.Write(add, data)
}

func (c *Cpu6502) Fetch() uint8 {
	data := c.Read(c.PC.Get())
	c.PC.Increment()
	return data
}

func (c *Cpu6502) Reset() {
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
func (c *Cpu6502) SET_OVERFLOW(cond bool) {
	c.Flag.Set(OVERFLOW, cond)
}

// Set CARRY if val is non zero
func (c *Cpu6502) SET_CARRY(val uint16) {
	c.Flag.Set(CARRY, val != 0)
}

// Set SIGN if val bit 7 is set
func (c *Cpu6502) SET_SIGN(val uint16) {
	c.Flag.Set(CARRY, val&0x80 != 0)
}

// Set Zero if val is zero
func (c *Cpu6502) SET_ZERO(val uint16) {
	c.Flag.Set(CARRY, val == 0)
}

// Set BREAk if val is non zero
func (c *Cpu6502) SET_BREAk(val uint16) {
	c.Flag.Set(CARRY, val != 0)
}

// Set INTERRUPT  if val is non zero
func (c *Cpu6502) SET_INTERRUPT(val uint16) {
	c.Flag.Set(INTERRUPT, val != 0)
}

// Set DECIMAL if val is non zero
func (c *Cpu6502) SET_DECIMAL(val uint16) {
	c.Flag.Set(DECIMAL, val != 0)
}

func (c *Cpu6502) PUSH(byte uint8) {
	c.Write(0x0100|c.SP.Get(), byte)
	c.SP.Decrement()
}

func (c *Cpu6502) PULL() uint8 {
	byte := c.Read(0x0100 | c.SP.Get())
	c.SP.Increment()
	return byte
}
