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

func TestZP(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0x85, 0xff, 0xa9, 0x00, 0x65, 0xff},
			A:    uint8(0xA2),
			SP:   uint8(0xFF),
			PC:   uint16(0x08),
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}

func TestZPX(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0x85, 0xff, 0xa2, 0x05, 0xa9, 0x00, 0x75, 0xfa},
			A:    uint8(0xA2),
			X:    uint8(0x05),
			SP:   uint8(0xFF),
			PC:   uint16(0x0A),
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}

func TestABS(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0x8d, 0x00, 0x10, 0xa9, 000, 0x6d, 0x00, 0x10},
			A:    uint8(0xA2),
			SP:   uint8(0xFF),
			PC:   uint16(0x0A),
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}
func TestABX(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0x8d, 0x05, 0x10, 0xa2, 0x05, 0xa9, 0x00, 0x7d, 0x00, 0x10},
			A:    uint8(0xA2),
			X:    uint8(0x05),
			SP:   uint8(0xFF),
			PC:   uint16(0x0C),
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}
func TestABY(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0x8d, 0x05, 0x10, 0xa0, 0x05, 0xa9, 0x00, 0x79, 0x00, 0x10},
			A:    uint8(0xA2),
			Y:    uint8(0x05),
			SP:   uint8(0xFF),
			PC:   uint16(0x0C),
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}
