package cpu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCPU6502_SetSR(t *testing.T) {
	type args struct {
		reg    SR
		enable bool
	}
	tests := []struct {
		name   string
		fields *CPU6502
		args   args
		exp    uint8
	}{
		{
			name:   "Set Carry",
			fields: &CPU6502{},
			args: args{
				reg:    CARRY,
				enable: true,
			},
			exp: uint8(CARRY),
		},
		{
			name:   "Set ZERO",
			fields: &CPU6502{},
			args: args{
				reg:    ZERO,
				enable: true,
			},
			exp: uint8(ZERO),
		},
		{
			name:   "Set SIGN",
			fields: &CPU6502{},
			args: args{
				reg:    SIGN,
				enable: true,
			},
			exp: uint8(SIGN),
		},
		{
			name: "Set SIGN & CARRY",
			fields: &CPU6502{
				SR: uint8(SIGN),
			},
			args: args{
				reg:    CARRY,
				enable: true,
			},
			exp: uint8(SIGN) | uint8(CARRY),
		},
		{
			name: "Clear Carry",
			fields: &CPU6502{
				SR: 0x01,
			},
			args: args{
				reg:    CARRY,
				enable: false,
			},
			exp: 0,
		},
		{
			name: "Clear ZERO",
			fields: &CPU6502{
				SR: 1 << 1,
			},
			args: args{
				reg:    ZERO,
				enable: false,
			},
			exp: 0,
		},
		{
			name: "Set SIGN & Clear CARRY",
			fields: &CPU6502{
				SR: uint8(SIGN) | uint8(CARRY),
			},
			args: args{
				reg:    CARRY,
				enable: false,
			},
			exp: uint8(SIGN),
		},
		{
			name: "Set SIGN & Clear CARRY (No Effect)",
			fields: &CPU6502{
				SR: uint8(SIGN),
			},
			args: args{
				reg:    CARRY,
				enable: false,
			},
			exp: uint8(SIGN),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.SetSR(tt.args.reg, tt.args.enable)
			assert.Equal(t, tt.exp, tt.fields.SR, "SetSP failed to set value")

			got := tt.fields.GetSR(tt.args.reg)
			assert.Equal(t, tt.args.enable, got, "GetSP failed to read value")
		})
	}
}
