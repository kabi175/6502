package model

type Opcode interface {
	Execute() uint8
	IsBreak() bool
	State() []byte
}

type OpcodeBuilder interface {
	Build(uint8) Opcode
}
