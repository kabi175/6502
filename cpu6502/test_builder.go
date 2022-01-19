package cpu6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kabi175/6502/bus"
	"github.com/kabi175/6502/cpu6502/debugger"
)

type Cputest struct {
	Prg  []uint8
	PC   uint16
	SP   uint8
	Flag uint8
	X    uint8
	Y    uint8
	A    uint8
	End  uint16
}

type Config struct {
	Prg    []uint8
	Breaks []uint16
	End    uint16
}

// Cpu builder package for internal testing
func NewCPUBuilder(c *Config) *CPU6502 {
	if c.End == 0 {
		c.End = uint16(len(c.Prg))
	}
	bus := bus.NewBus16(c.Prg)
	deb := debugger.NewDebugger(c.Breaks, c.End)
	cpu := newCPU(bus)
	cpu.AddEvent(deb)
	return cpu
}

func ProgramTest(t *testing.T, tests []Cputest) {
	for id, test := range tests {
		testID := fmt.Sprintf("SubTest_%v", id)
		t.Run(testID, func(t *testing.T) {
			cpu := NewCPUBuilder(&Config{Prg: test.Prg, End: test.End})
			cpu.Execute()
			assert.Equalf(t, test.A, cpu.A.Get(), "Error A Register")
			assert.Equalf(t, test.X, cpu.X.Get(), "Error X Register")
			assert.Equalf(t, test.Y, cpu.Y.Get(), "Error Y Register")
			assert.Equalf(t, test.SP, uint8(cpu.SP.Get()), "Error SP")
			assert.Equalf(t, test.PC, cpu.PC.Get(), "Error PC")
			assert.Equalf(t, test.Flag, cpu.Flag.Byte(), "Error Flag Register:%b", cpu.Flag.Byte())
		})
	}
}

func OpcodeTest(t *testing.T, hex uint8, ins INSTRUCTION, mode ADDRMODE) {
	got := NewOpcode(hex)
	assert.Equal(t, ins, got.Ins, "Invalid Instruction")
	assert.Equal(t, mode, got.Mode, "Invalid Addressing Mode")
}
