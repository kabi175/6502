package main

import (
	"github.com/kabi175/6502/bus"
	"github.com/kabi175/6502/cpu6502"
	"github.com/kabi175/6502/cpu6502/opcode"
	"github.com/kabi175/6502/graphic/debugger"
)

func main() {
	ram := []uint8{0xa9, 0x01, 0x8d, 0x00, 0x02, 0xa9, 0x05, 0x8d, 0x01, 0x02, 0xa9, 0x08, 0x8d, 0x02, 0x02}
	//ram := []uint8{0xa9, 0xc0, 0xaa, 0xe8, 0x69, 0xc4, 0x00}
	bus16 := bus.NewBus16(ram)
	deb := debugger.NewDebugger()
	cpu := cpu6502.New(bus16, opcode.NewOpcode, deb)
	cpu.Execute(deb.Quit)
	defer deb.Close()
}
