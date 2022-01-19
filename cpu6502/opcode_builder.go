package cpu6502

import (
	"github.com/kabi175/6502/model"
)

func NewOpcodeBuilder(cpu *CPU6502) *OpcodeBuilder {
	return &OpcodeBuilder{
		cpu: cpu,
	}
}

type OpcodeBuilder struct {
	cpu *CPU6502
}

func (b *OpcodeBuilder) Build(opcodeHex uint8) model.Opcode {
	opcode := NewOpcode(opcodeHex)
	opcode.cpu = b.cpu
	return opcode
}
