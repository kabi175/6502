package opcode_test

import (
	"testing"

	"github.com/kabi175/6502/cpu6502/test"
)

func TestInterrupt(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0x78},
			PC:   uint16(0x01),
			SP:   uint8(0xff),
			Flag: 0b00100100,
		},
		{
			Prg:  []uint8{0x58},
			PC:   uint16(0x01),
			SP:   uint8(0xff),
			Flag: 0b00100000,
		},
		{
			Prg:  []uint8{0x78, 0x58},
			PC:   uint16(0x02),
			SP:   uint8(0xff),
			Flag: 0b00100000,
		},
	}
	test.ProgramTest(t, tests)
}

func TestDecimal(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xf8},
			PC:   uint16(0x01),
			SP:   uint8(0xff),
			Flag: 0b00101000,
		},
		{
			Prg:  []uint8{0xd8},
			PC:   uint16(0x01),
			SP:   uint8(0xff),
			Flag: 0b00100000,
		},
		{
			Prg:  []uint8{0xf8, 0xd8},
			PC:   uint16(0x02),
			SP:   uint8(0xff),
			Flag: 0b00100000,
		},
	}
	test.ProgramTest(t, tests)
}

func TestCarry(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0x38},
			PC:   uint16(0x01),
			SP:   uint8(0xff),
			Flag: 0b00100001,
		},
		{
			Prg:  []uint8{0x18},
			PC:   uint16(0x01),
			SP:   uint8(0xff),
			Flag: 0b00100000,
		},
		{
			Prg:  []uint8{0x38, 0x18},
			PC:   uint16(0x02),
			SP:   uint8(0xff),
			Flag: 0b00100000,
		},
	}
	test.ProgramTest(t, tests)
}
