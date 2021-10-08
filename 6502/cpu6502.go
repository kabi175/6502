package cpu6502

type cpu6502 struct {
	operand uint8
	Opcodes map[uint16]Instruction
	PC      PC16
	SP      SP8
	A       GP8
	X       GP8
	Y       GP8
	Bus     Bus16
}

//	Constructor function to create CPU6502
func New(bus Bus16) *cpu6502 {
	return &cpu6502{
		Opcodes: load(),
		PC:      NewPC(),
		SP:      NewSP8(),
		A:       NewGP8(),
		X:       NewGP8(),
		Y:       NewGP8(),
		Bus:     bus,
	}
}

//	check chan for new message or chan close
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
func (c *cpu6502) Execute(close chan bool) {
	for !isClosed(close) {
		opcode := c.fetch()
		ins := c.Opcodes[uint16(opcode)]
		cycle := ins.Cycle
		cycle += ins.Mode(c)
		cycle += ins.Ins(c)
	}
}

func (c *cpu6502) read(add uint16) uint8 {
	return c.Bus.Read(add)
}

func (c *cpu6502) write(add uint16, data uint8) {
	c.Bus.Write(add, data)
}

func (c *cpu6502) fetch() uint8 {
	data := c.read(c.PC.Get())
	c.PC.Increment()
	return data
}
