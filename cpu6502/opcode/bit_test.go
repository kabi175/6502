package opcode_test

import (
	"testing"

	"github.com/kabi175/6502/cpu6502/test"
)

func TestBIT(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xf3, 0x85, 0xf0, 0x24, 0xf0},
			A:    0xF3,
			SP:   0xFF,
			PC:   0x06,
			Flag: 0b11100000,
		},
		{
			Prg:  []uint8{0xa9, 0xa3, 0x85, 0xf0, 0xa9, 0x00, 0x24, 0xf0},
			A:    0x0,
			SP:   0xFF,
			PC:   0x08,
			Flag: 0b10100010,
		},
	}
	test.ProgramTest(t, tests)
}
