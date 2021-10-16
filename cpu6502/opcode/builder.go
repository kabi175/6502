package opcode

import (
	"github.com/kabi175/6502/cpu6502"
	"github.com/kabi175/6502/model"
)

func NewOpcodeBuilder(cpu *cpu6502.Cpu6502) *OpcodeBuilder {
	return &OpcodeBuilder{
		cpu: cpu,
	}
}

type OpcodeBuilder struct {
	cpu *cpu6502.Cpu6502
}

func (b *OpcodeBuilder) Build(opcodeHex uint8) model.Opcode {
	opcode := NewOpcode(opcodeHex)
	opcode.cpu = b.cpu
	return opcode
}
