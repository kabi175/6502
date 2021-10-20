package opcode_test

import (
	"testing"

	"github.com/kabi175/6502/cpu6502/test"
)

func TestJMP(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0x4c, 0x23, 0x23},
			PC:   uint16(0x2323),
			SP:   uint8(0xff),
			Flag: 0b00100000,
			End:  uint16(0x2323),
		},
	}
	test.ProgramTest(t, tests)
}
