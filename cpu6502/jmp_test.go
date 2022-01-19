package cpu6502

import (
	"testing"
)

func TestJMP(t *testing.T) {
	tests := []Cputest{
		{
			Prg:  []uint8{0x4c, 0x23, 0x23},
			PC:   uint16(0x2323),
			SP:   uint8(0xff),
			Flag: 0b00100000,
			End:  uint16(0x2323),
		},
	}
	ProgramTest(t, tests)
}
