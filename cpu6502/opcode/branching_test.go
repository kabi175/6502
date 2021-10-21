package opcode_test

import (
	"testing"

	"github.com/kabi175/6502/cpu6502/test"
)

func TestBCC(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0x38, 0x69, 0x01, 0x90, 0x03, 0x4c, 0x00, 0xf0},
			A:    uint8(0x02),
			SP:   uint8(0xff),
			PC:   uint16(0xf000),
			Flag: 0b00100000,
			End:  uint16(0xf000),
		},
	}
	test.ProgramTest(t, tests)
}

func TestBCS(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xff, 0x69, 0x01, 0xb0, 0x03, 0x4c, 0x02, 0x00, 0x85, 0xf0},
			A:    uint8(0x00),
			PC:   uint16(0x09),
			SP:   uint8(0xff),
			Flag: 0b00100011,
			End:  uint16(0x09),
		},
	}
	test.ProgramTest(t, tests)
}

func TestBEQ(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xff, 0x69, 0x01, 0xf0, 0x03, 0x4c, 0x02, 0x00, 0x85, 0xf0},
			A:    uint8(0x00),
			PC:   uint16(0x09),
			SP:   uint8(0xff),
			Flag: 0b00100011,
			End:  uint16(0x09),
		},
	}
	test.ProgramTest(t, tests)
}