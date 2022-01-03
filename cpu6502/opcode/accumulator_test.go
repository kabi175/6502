package opcode_test

import (
	"testing"

	"github.com/kabi175/6502/cpu6502/opcode"
	"github.com/kabi175/6502/cpu6502/test"
)

func TestADC(t *testing.T) {
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
		// Zero Page
		{
			Prg:  []uint8{0xa9, 0xa2, 0x85, 0xff, 0xa9, 0x00, 0x65, 0xff},
			A:    uint8(0xA2),
			SP:   uint8(0xFF),
			PC:   uint16(0x08),
			Flag: 0b10100000,
		},
		// Zero Page X
		{
			Prg:  []uint8{0xa9, 0xa2, 0x85, 0xff, 0xa2, 0x05, 0xa9, 0x00, 0x75, 0xfa},
			A:    uint8(0xA2),
			X:    uint8(0x05),
			SP:   uint8(0xFF),
			PC:   uint16(0x0A),
			Flag: 0b10100000,
		},
		{
			Prg:  []uint8{0xa9, 0xa2, 0x8d, 0x00, 0x10, 0xa9, 000, 0x6d, 0x00, 0x10},
			A:    uint8(0xA2),
			SP:   uint8(0xFF),
			PC:   uint16(0x0A),
			Flag: 0b10100000,
		},
		{
			Prg:  []uint8{0xa9, 0xa2, 0x8d, 0x05, 0x10, 0xa2, 0x05, 0xa9, 0x00, 0x7d, 0x00, 0x10},
			A:    uint8(0xA2),
			X:    uint8(0x05),
			SP:   uint8(0xFF),
			PC:   uint16(0x0C),
			Flag: 0b10100000,
		},
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

func TestAND(t *testing.T) {
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
		{
			Prg:  []uint8{0xa9, 0xa4, 0x85, 0xf0, 0xa9, 0xc6, 0x25, 0xf0},
			A:    uint8(0x84),
			SP:   uint8(0xff),
			PC:   uint16(0x08),
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)

	// Zero page
	test.OpcodeTest(t, 0x25, opcode.AND, opcode.ZPA)
	test.OpcodeTest(t, 0x35, opcode.AND, opcode.ZPX)

	// Absolute Mode
	test.OpcodeTest(t, 0x2D, opcode.AND, opcode.ABS)
	test.OpcodeTest(t, 0x3D, opcode.AND, opcode.ABX)
	test.OpcodeTest(t, 0x39, opcode.AND, opcode.ABY)
	// Relative Indexed
	test.OpcodeTest(t, 0x21, opcode.AND, opcode.IDX)
	test.OpcodeTest(t, 0x31, opcode.AND, opcode.IDY)
}

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

func TestCMP(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa4, 0x85, 0xf0, 0xa9, 0xc6, 0xc5, 0xf0},
			A:    0xC6,
			PC:   0x08,
			SP:   0xFF,
			Flag: 0b00100001,
		},
	}
	test.ProgramTest(t, tests)
}

func TestDEC(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa4, 0x85, 0xf0, 0xc6, 0xf0, 0xa5, 0xf0},
			A:    0xA3,
			PC:   0x08,
			SP:   0xFF,
			Flag: 0b10100000,
		},
		{
			Prg:  []uint8{0xa9, 0x00, 0x85, 0xf0, 0xc6, 0xf0, 0xa5, 0xf0},
			A:    0xFF,
			PC:   0x08,
			SP:   0xFF,
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}
func TestEOR(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa4, 0x85, 0xf0, 0xa9, 0x45, 0x45, 0xf0},
			A:    0xE1,
			PC:   0x08,
			SP:   0xFF,
			Flag: 0b10100000,
		},
	}
	test.ProgramTest(t, tests)
}
func TestLSR(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0x4a},
			A:    0x51,
			PC:   0x03,
			SP:   0xFF,
			Flag: 0b00100000,
		},
		{
			Prg:  []uint8{0xa9, 0xa2, 0x85, 0xf0, 0x46, 0xf0, 0xa5, 0xf0},
			A:    0x51,
			PC:   0x08,
			SP:   0xFF,
			Flag: 0b00100000,
		},
	}
	test.ProgramTest(t, tests)
}

func TestORA(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0x85, 0xf0, 0x46, 0xf0, 0x05, 0xf0},
			A:    0xF3,
			PC:   0x08,
			SP:   0xFF,
			Flag: 0b10100000,
		},
		{
			Prg:  []uint8{0xa9, 0x00, 0x85, 0xf0, 0x46, 0xf0, 0x05, 0xf0},
			A:    0x00,
			PC:   0x08,
			SP:   0xFF,
			Flag: 0b00100010,
		},
	}
	test.ProgramTest(t, tests)
}

func TestROL(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0x2a},
			A:    0x44,
			PC:   0x03,
			SP:   0xFF,
			Flag: 0b00100001,
		},
	}
	test.ProgramTest(t, tests)
}

func TestROR(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0x6a},
			A:    0x51,
			PC:   0x03,
			SP:   0xFF,
			Flag: 0b00100000,
		},
	}
	test.ProgramTest(t, tests)
}

func TestSBC(t *testing.T) {
	tests := []test.Cputest{
		{
			Prg:  []uint8{0xa9, 0xa2, 0xe9, 0x34},
			A:    0x6d,
			PC:   0x04,
			SP:   0xFF,
			Flag: 0b01100001,
		},
		//  Decimal Mode
		{
			Prg:  []uint8{0xf8, 0xa9, 0xa2, 0xe9, 0x34},
			A:    0x67,
			PC:   0x05,
			SP:   0xFF,
			Flag: 0b01101001,
		},
	}
	test.ProgramTest(t, tests)
}
