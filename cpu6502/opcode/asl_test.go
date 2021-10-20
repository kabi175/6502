package opcode_test

import (
	"testing"

	"github.com/kabi175/6502/cpu6502/test"
)

func Test_ASL(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0x34, 0x0a},
			A:    uint8(0x68),
			PC:   uint16(0x03),
			SP:   uint8(0xff),
			Flag: 0b00100000,
		},
		{
			Prg:  []uint8{0xa9, 0x43, 0x85, 0xf0, 0x06, 0xf0, 0xa5, 0xf0},
			A:    uint8(0x86),
			PC:   uint16(0x08),
			SP:   uint8(0xff),
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}
