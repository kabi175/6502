package model

// Opcode interface
// Capable of changing the CPU state depending on the  instruction.
type Opcode interface {
	Execute() uint8
	IsBreak() bool
	State() []byte
}

// Opcode Builder Interface
type OpcodeBuilder interface {
	Build(uint8) Opcode
}
