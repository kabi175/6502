package cpu

import (
	"testing"

	"github.com/kabi175/6502/model"
	"github.com/stretchr/testify/assert"
)

func uint8P(in uint8) *uint8 {
	return &in
}
func uint16P(in uint16) *uint16 {
	return &in
}

type bus map[uint16]uint8

// Read implements model.Bus16.
func (b bus) Read(addr uint16) uint8 {
	v, ok := b[addr]
	if !ok {
		return 0
	}
	return v
}

// Write implements model.Bus16.
func (b bus) Write(addr uint16, data uint8) {
	b[addr] = data
}

var _ model.Bus16 = make(bus, 0)

func TestACC(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
	}{
		{
			name:      "Simple Accmulator",
			args:      args{&CPU6502{AC: 0x02}},
			wantOpr:   uint8P(0x02),
			wantAdrr:  nil,
			wantCycle: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := ACC(tt.args.cpu)
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}

func TestIMP(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
	}{
		{
			name:      "implied Addressing, no operand",
			args:      args{&CPU6502{}},
			wantOpr:   nil,
			wantAdrr:  nil,
			wantCycle: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := IMP(tt.args.cpu)
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}

func TestIMM(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
		wantPC    uint16
	}{
		{
			name: "Immediate Addressing T1",
			args: args{&CPU6502{
				PC:  0x01,
				bus: bus{0x01: 0x03},
			}},
			wantOpr:   uint8P(0x03),
			wantAdrr:  nil,
			wantCycle: 1,
			wantPC:    0x02,
		},
		{
			name: "Immediate Addressing T2",
			args: args{&CPU6502{
				PC:  0xFFFF,
				bus: bus{0xFFFF: 0xF2},
			}},
			wantOpr:   uint8P(0xF2),
			wantAdrr:  nil,
			wantCycle: 1,
			wantPC:    0x00,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := IMM(tt.args.cpu)
			assert.Equal(t, tt.wantPC, tt.args.cpu.PC, "Got invalid PC")
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}

func TestABS(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
		wantPC    uint16
	}{
		{
			name: "absolute samepage1",
			args: args{&CPU6502{
				PC:  0x01,
				X:   3,
				Y:   2,
				bus: bus{0x01: 0xFF, 0x02: 0x00, 0xFF: 0x04},
			}},
			wantOpr:   uint8P(0x04),
			wantAdrr:  uint16P(0x00FF),
			wantCycle: 3,
			wantPC:    0x03,
		},
		{
			name: "absolute samepage2",
			args: args{&CPU6502{
				PC:  0x44,
				X:   0xA3,
				Y:   0xC2,
				bus: bus{0x44: 0xFF, 0x45: 0x0F, 0x0FFF: 0x04},
			}},
			wantOpr:   uint8P(0x04),
			wantAdrr:  uint16P(0x0FFF),
			wantCycle: 3,
			wantPC:    0x46,
		},
		{
			name: "absolute diffpage",
			args: args{&CPU6502{
				X:   0xE3,
				Y:   0xD2,
				PC:  0x5545,
				bus: bus{0x5545: 0xF4, 0x5546: 0x34, 0x34F4: 0x5F},
			}},
			wantOpr:   uint8P(0x5F),
			wantAdrr:  uint16P(0x34F4),
			wantCycle: 3,
			wantPC:    0x5547,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := ABS(tt.args.cpu)
			assert.Equal(t, tt.wantPC, tt.args.cpu.PC, "Got invalid PC")
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}

func TestABS_X(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
		wantPC    uint16
	}{
		{
			name: "absolute samepage",
			args: args{&CPU6502{
				PC:  0x01,
				X:   0x03,
				Y:   0x02,
				bus: bus{0x01: 0xF0, 0x02: 0x00, 0xF3: 0x04},
			}},
			wantOpr:   uint8P(0x04),
			wantAdrr:  uint16P(0xF3),
			wantCycle: 3,
			wantPC:    0x03,
		},
		{
			name: "absolute crosspage",
			args: args{&CPU6502{
				PC:  0x44,
				X:   0xA7,
				Y:   0xC2,
				bus: bus{0x44: 0xFF, 0x45: 0x0F, 0x10a6: 0x04},
			}},
			wantOpr:   uint8P(0x04),
			wantAdrr:  uint16P(0x10a6),
			wantCycle: 4,
			wantPC:    0x46,
		},
		{
			name: "absolute max-offset",
			args: args{&CPU6502{
				X:   0xFF,
				Y:   0xD2,
				PC:  0x5545,
				bus: bus{0x5545: 0xF4, 0x5546: 0x34, 0x35F3: 0x5F},
			}},
			wantOpr:   uint8P(0x5F),
			wantAdrr:  uint16P(0x35F3),
			wantCycle: 4,
			wantPC:    0x5547,
		},
		{
			name: "absolute 0-offset",
			args: args{&CPU6502{
				X:   0x0,
				Y:   0xD2,
				PC:  0x5545,
				bus: bus{0x5545: 0xF4, 0x5546: 0x34, 0x34F4: 0x5F},
			}},
			wantOpr:   uint8P(0x5F),
			wantAdrr:  uint16P(0x34F4),
			wantCycle: 3,
			wantPC:    0x5547,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := ABS_X(tt.args.cpu)
			assert.Equal(t, tt.wantPC, tt.args.cpu.PC, "Got invalid PC")
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}

func TestABS_Y(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
		wantPC    uint16
	}{
		{
			name: "samepage",
			args: args{&CPU6502{
				PC:  0x01,
				Y:   0x03,
				X:   0x02,
				bus: bus{0x01: 0xF0, 0x02: 0x00, 0xF3: 0x04},
			}},
			wantOpr:   uint8P(0x04),
			wantAdrr:  uint16P(0xF3),
			wantCycle: 3,
			wantPC:    0x03,
		},
		{
			name: "crosspage",
			args: args{&CPU6502{
				PC:  0x44,
				Y:   0xA7,
				X:   0xC2,
				bus: bus{0x44: 0xFF, 0x45: 0x0F, 0x10a6: 0x04},
			}},
			wantOpr:   uint8P(0x04),
			wantAdrr:  uint16P(0x10a6),
			wantCycle: 4,
			wantPC:    0x46,
		},
		{
			name: "max-offset",
			args: args{&CPU6502{
				Y:   0xFF,
				X:   0xD2,
				PC:  0x5545,
				bus: bus{0x5545: 0xF4, 0x5546: 0x34, 0x35F3: 0x5F},
			}},
			wantOpr:   uint8P(0x5F),
			wantAdrr:  uint16P(0x35F3),
			wantCycle: 4,
			wantPC:    0x5547,
		},
		{
			name: "0-offset",
			args: args{&CPU6502{
				Y:   0x0,
				X:   0xD2,
				PC:  0x5545,
				bus: bus{0x5545: 0xF4, 0x5546: 0x34, 0x34F4: 0x5F},
			}},
			wantOpr:   uint8P(0x5F),
			wantAdrr:  uint16P(0x34F4),
			wantCycle: 3,
			wantPC:    0x5547,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := ABS_Y(tt.args.cpu)
			assert.Equal(t, tt.wantPC, tt.args.cpu.PC, "Got invalid PC")
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}

func TestZP(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
		wantPC    uint16
	}{
		{
			name: "T1",
			args: args{&CPU6502{
				PC:  0x01,
				Y:   0x03,
				X:   0x02,
				bus: bus{0x01: 0xF0, 0xF0: 0x04},
			}},
			wantOpr:   uint8P(0x04),
			wantAdrr:  uint16P(0xF0),
			wantCycle: 2,
			wantPC:    0x02,
		},
		{
			name: "T2",
			args: args{&CPU6502{
				PC:  0x6244,
				Y:   0xA7,
				X:   0xC2,
				bus: bus{0x6244: 0xFF, 0xFF: 0x5F},
			}},
			wantOpr:   uint8P(0x5F),
			wantAdrr:  uint16P(0xFF),
			wantCycle: 2,
			wantPC:    0x6245,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := ZP(tt.args.cpu)
			assert.Equal(t, tt.wantPC, tt.args.cpu.PC, "Got invalid PC")
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}

func TestZP_X(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
		wantPC    uint16
	}{
		{
			name: "no overflow",
			args: args{&CPU6502{
				PC:  0x01,
				Y:   0x03,
				X:   0x02,
				bus: bus{0x01: 0xF0, 0xF2: 0x04},
			}},
			wantOpr:   uint8P(0x04),
			wantAdrr:  uint16P(0xF2),
			wantCycle: 2,
			wantPC:    0x02,
		},
		{
			name: "overflow",
			args: args{&CPU6502{
				PC:  0x6244,
				Y:   0xA7,
				X:   0xC2,
				bus: bus{0x6244: 0xFF, 0xC1: 0x5F},
			}},
			wantOpr:   uint8P(0x5F),
			wantAdrr:  uint16P(0xC1),
			wantCycle: 2,
			wantPC:    0x6245,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := ZP_X(tt.args.cpu)
			assert.Equal(t, tt.wantPC, tt.args.cpu.PC, "Got invalid PC")
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}

func TestZP_Y(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
		wantPC    uint16
	}{
		{
			name: "no overflow",
			args: args{&CPU6502{
				PC:  0x01,
				X:   0x03,
				Y:   0x02,
				bus: bus{0x01: 0xF0, 0xF2: 0x04},
			}},
			wantOpr:   uint8P(0x04),
			wantAdrr:  uint16P(0xF2),
			wantCycle: 2,
			wantPC:    0x02,
		},
		{
			name: "overflow",
			args: args{&CPU6502{
				PC:  0x6244,
				X:   0xA7,
				Y:   0xC2,
				bus: bus{0x6244: 0xFF, 0xC1: 0x5F},
			}},
			wantOpr:   uint8P(0x5F),
			wantAdrr:  uint16P(0xC1),
			wantCycle: 2,
			wantPC:    0x6245,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := ZP_Y(tt.args.cpu)
			assert.Equal(t, tt.wantPC, tt.args.cpu.PC, "Got invalid PC")
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}

func TestIND(t *testing.T) {
	type args struct {
		cpu *CPU6502
	}
	tests := []struct {
		name      string
		args      args
		wantOpr   *uint8
		wantAdrr  *uint16
		wantCycle uint8
		wantPC    uint16
	}{
		{
			name: "T1",
			args: args{&CPU6502{
				PC:  0x01,
				X:   0x03,
				Y:   0x02,
				bus: bus{0x01: 0x82, 0x02: 0xFF, 0xFF82: 0xC4, 0xFF83: 0x80},
			}},
			wantOpr:   nil,
			wantAdrr:  uint16P(0x80C4),
			wantCycle: 4,
			wantPC:    0x03,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOpr, gotAdrr, gotCycle := IND(tt.args.cpu)
			assert.Equal(t, tt.wantPC, tt.args.cpu.PC, "Got invalid PC")
			assert.Equal(t, tt.wantOpr, gotOpr, "Got invalid Operator")
			assert.Equal(t, tt.wantAdrr, gotAdrr, "Got invalid Address")
			assert.Equal(t, tt.wantCycle, gotCycle, "Got invalid Cycle")
		})
	}
}
