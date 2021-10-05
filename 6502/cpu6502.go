package cpu6502

type Bus16 interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}

type cpu6502 struct {
	Opcodes map[uint8]Instruction
	PC      PC16
	SP      SP8
	Bus     Bus16
}

func New(bus Bus16) *cpu6502 {
	return &cpu6502{
		Opcodes: load(),
		PC:      NewPC(),
		SP:      NewSP8(),
		Bus:     bus,
	}
}

func (c *cpu6502) Execute() {}

func (c *cpu6502) read(add uint16) uint8 {
	return c.Bus.Read(add)
}

func (c *cpu6502) write(add uint16, data uint8) {
	c.Bus.Write(add, data)
}
