package opcode_test

import (
	"testing"

	"github.com/kabi175/6502/cpu6502/test"
)

func TestAND_IMM(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0x2a, 0x29, 0x32},
			SP:   uint8(0xff),
			PC:   uint16(0x04),
			A:    uint8(0x22),
			Flag: 0b00100000,
		},
		{
			Prg:  []uint8{0xa9, 0x2a, 0x29, 0x00},
			SP:   uint8(0xff),
			PC:   uint16(0x04),
			A:    uint8(0x00),
			Flag: 0b00100010,
		},
		{
			Prg:  []uint8{0xa9, 0xff, 0x29, 0xff},
			SP:   uint8(0xff),
			PC:   uint16(0x04),
			A:    uint8(0xff),
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}

func TestADC_ZP(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa4, 0x85, 0xf0, 0xa9, 0xc6, 0x25, 0xf0},
			A:    uint8(0x84),
			SP:   uint8(0xff),
			PC:   uint16(0x08),
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}
