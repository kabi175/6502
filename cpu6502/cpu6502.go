package cpu6502

type Cpu6502 struct {
	Operand uint8
	Addr    uint16
	Opcodes map[uint16]Opcode
	PC      PC16
	SP      SP8
	Flag    FlagRegister
	A       GP8
	X       GP8
	Y       GP8
	Bus     Bus16
}

//	Constructor function to create CPU6502
func New(bus Bus16, opcodes map[uint16]Opcode) *Cpu6502 {
	return &Cpu6502{
		Opcodes: opcodes,
		PC:      NewPC(),
		SP:      NewSP8(),
		Flag:    NewFlagRegister(),
		A:       NewGP8(),
		X:       NewGP8(),
		Y:       NewGP8(),
		Bus:     bus,
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
		opcode := c.Fetch()
		c.Opcodes[uint16(opcode)].Execute(c)
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
