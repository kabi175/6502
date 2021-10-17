package opcode_test

import (
	"testing"

	"github.com/kabi175/6502/cpu6502/test"
)

func TestIMM(t *testing.T) {
	tests := []test.Cputest{
		// NON Decial Mode Test
		{
			Prg:  []uint8{0x69, 0x10},
			PC:   uint16(0x02),
			SP:   uint8(0xFF),
			Flag: uint8(0b00100000),
			X:    uint8(0x00),
			Y:    uint8(0x00),
			A:    uint8(0x10),
		},
		{
			Prg:  []uint8{0x38, 0x69, 0x10},
			PC:   uint16(0x03),
			SP:   uint8(0xFF),
			Flag: uint8(0b00100000),
			A:    uint8(0x11),
		},
		{
			Prg:  []uint8{0x38, 0x69, 0xff},
			PC:   uint16(0x03),
			SP:   uint8(0xFF),
			Flag: uint8(0b00100011),
			A:    uint8(0x00),
		},
		// Decimal Mode Test
	}
	test.ProgramTest(t, tests)
}

func TestDecimalMode(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xf8, 0x38, 0x69, 0x10},
			PC:   uint16(0x04),
			SP:   uint8(0xFF),
			A:    uint8(0x11),
			Flag: uint8(0b00101000),
		},
		{
			Prg:  []uint8{0xf8, 0xa9, 0x10, 0x69, 0x10},
			PC:   uint16(0x05),
			SP:   uint8(0xFF),
			A:    uint8(0x20),
			Flag: uint8(0b00101000),
		},
		{
			Prg:  []uint8{0xf8, 0xa9, 0x05, 0x69, 0x05},
			PC:   uint16(0x05),
			SP:   uint8(0xFF),
			A:    uint8(0x10),
			Flag: uint8(0b00101000),
		},
		{
			Prg:  []uint8{0xf8, 0xa9, 0x50, 0x69, 0x50},
			PC:   uint16(0x05),
			SP:   uint8(0xFF),
			A:    uint8(0x00),
			Flag: uint8(0b01101011),
		},
	}

	test.ProgramTest(t, tests)
}
