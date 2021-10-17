package opcode_test

import (
	"testing"

	"github.com/kabi175/6502/cpu6502/test"
)

func TestIMM(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0x10},
			PC:   uint16(0x02),
			SP:   uint8(0xFF),
			Flag: uint8(0b00100000),
			X:    uint8(0x00),
			Y:    uint8(0x00),
			A:    uint8(0x10),
		},
	}
	test.ProgramTest(t, tests)
}
