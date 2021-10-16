package main

import (
	"github.com/kabi175/6502/bus"
	"github.com/kabi175/6502/cpu6502"
	"github.com/kabi175/6502/cpu6502/opcode"
)

func main() {
	ram := []uint8{0xa9, 0x01, 0x8d, 0x00, 0x02, 0xa9, 0x05, 0x8d, 0x01, 0x02, 0xa9, 0x08, 0x8d, 0x02, 0x02}
	bus16 := bus.NewBus16(ram)
	cpu := cpu6502.NewCpu(bus16, nil)
	builder := opcode.NewOpcodeBuilder(cpu)
	cpu.AttachOpcodeBuilder(builder)
}
