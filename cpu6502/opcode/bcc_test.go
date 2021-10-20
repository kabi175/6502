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
