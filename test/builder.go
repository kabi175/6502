package test

import (
	"github.com/kabi175/6502/bus"
	"github.com/kabi175/6502/cpu6502"
	"github.com/kabi175/6502/cpu6502/debugger"
	"github.com/kabi175/6502/cpu6502/opcode"
)

type Config struct {
	Prg    []uint8
	Breaks []uint16
	End    uint16
}

// Cpu builder package for internal testing
func NewCpuBuilder(c *Config) *cpu6502.Cpu6502 {
	if c.End == 0 {
		c.End = uint16(len(c.Prg) - 1)
	}
	bus := bus.NewBus16(c.Prg)
	deb := debugger.NewDebugger(c.Breaks, c.End)
	cpu := cpu6502.NewCpu(bus, deb)
	opcodeBuilder := opcode.NewOpcodeBuilder(cpu)
	cpu.AttachOpcodeBuilder(opcodeBuilder)
	return cpu
}
