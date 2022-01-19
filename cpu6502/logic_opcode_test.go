package cpu6502

import "testing"

func Test_opcode_ADC(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.ADC(tt.args.c); got != tt.want {
				t.Errorf("opcode.ADC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_AND(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.AND(tt.args.c); got != tt.want {
				t.Errorf("AND() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_ASL(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.ASL(tt.args.c); got != tt.want {
				t.Errorf("opcode.ASL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_BIT(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.BIT(tt.args.c); got != tt.want {
				t.Errorf("opcode.BIT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_CMP(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.CMP(tt.args.c); got != tt.want {
				t.Errorf("opcode.CMP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_DEC(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.DEC(tt.args.c); got != tt.want {
				t.Errorf("opcode.DEC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_EOR(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.EOR(tt.args.c); got != tt.want {
				t.Errorf("opcode.EOR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_LSR(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.LSR(tt.args.c); got != tt.want {
				t.Errorf("opcode.LSR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_ORA(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.ORA(tt.args.c); got != tt.want {
				t.Errorf("opcode.ORA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_ROL(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.ROL(tt.args.c); got != tt.want {
				t.Errorf("opcode.ROL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_ROR(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.ROR(tt.args.c); got != tt.want {
				t.Errorf("opcode.ROR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_opcode_SBC(t *testing.T) {
	type fields struct {
		Cycle uint8
		Mode  ADDRMODE
		Ins   INSTRUCTION
		cpu   *CPU6502
	}
	type args struct {
		c *CPU6502
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opcode{
				Cycle: tt.fields.Cycle,
				Mode:  tt.fields.Mode,
				Ins:   tt.fields.Ins,
				cpu:   tt.fields.cpu,
			}
			if got := o.SBC(tt.args.c); got != tt.want {
				t.Errorf("opcode.SBC() = %v, want %v", got, tt.want)
			}
		})
	}
}
