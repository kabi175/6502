package opcode

import (
	"fmt"
)

func NewOpcode(code uint8) *opcode {
	switch code {
	case 0x00:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   BRK,
		}
	case 0x01:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   ORA,
		}
	case 0x05:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   ORA,
		}
	case 0x06:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   ASL,
		}
	case 0x08:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   PHP,
		}
	case 0x09:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   ORA,
		}
	case 0x0A:
		return &opcode{
			Cycle: 0,
			Mode:  ACC,
			Ins:   ASL,
		}

	case 0x0D:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   ORA,
		}
	case 0x0E:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   ASL,
		}
	case 0x10:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   BPL,
		}
	case 0x11:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   ORA,
		}
	case 0x15:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   ORA,
		}
	case 0x16:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   ASL,
		}
	case 0x18:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   CLC,
		}
	case 0x19:
		return &opcode{
			Cycle: 0,
			Mode:  ABY,
			Ins:   ORA,
		}
	case 0x1D:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   ORA,
		}
	case 0x1E:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   ASL,
		}
	case 0x20:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   JSR,
		}
	case 0x21:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   AND,
		}
	case 0x24:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   BIT,
		}
	case 0x25:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   AND,
		}
	case 0x26:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   ROL,
		}
	case 0x28:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   PLP,
		}
	case 0x29:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   AND,
		}
	case 0x2A:
		return &opcode{
			Cycle: 0,
			Mode:  ACC,
			Ins:   ROL,
		}
	case 0x2C:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   BIT,
		}
	case 0x2D:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   AND,
		}
	case 0x2E:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   ROL,
		}
	case 0x30:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   BMI,
		}
	case 0x31:
		return &opcode{
			Cycle: 0,
			Mode:  IDY,
			Ins:   AND,
		}
	case 0x35:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   AND,
		}
	case 0x36:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   ROL,
		}
	case 0x38:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   SEC,
		}
	case 0x39:
		return &opcode{
			Cycle: 0,
			Mode:  ABY,
			Ins:   AND,
		}
	case 0x3D:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   AND,
		}
	case 0x3E:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   ROL,
		}
	case 0x40:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   RTI,
		}
	case 0x41:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   EOR,
		}
	case 0x45:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   EOR,
		}
	case 0x46:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   LSR,
		}
	case 0x48:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   PHA,
		}
	case 0x49:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   EOR,
		}
	case 0x4A:
		return &opcode{
			Cycle: 0,
			Mode:  ACC,
			Ins:   LSR,
		}
	case 0x4C:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   JMP,
		}
	case 0x4D:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   EOR,
		}
	case 0x4E:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   LSR,
		}
	case 0x50:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   BVC,
		}
	case 0x51:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   EOR,
		}
	case 0x55:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   EOR,
		}
	case 0x56:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   LSR,
		}
	case 0x58:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   CLI,
		}
	case 0x59:
		return &opcode{
			Cycle: 0,
			Mode:  ABY,
			Ins:   EOR,
		}
	case 0x5D:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   EOR,
		}
	case 0x5E:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   LSR,
		}
	case 0x60:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   RTS,
		}
	case 0x61:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   ADC,
		}
	case 0x65:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   ADC,
		}
	case 0x66:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   ROR,
		}
	case 0x68:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   PLA,
		}
	case 0x69:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   ADC,
		}
	case 0x6A:
		return &opcode{
			Cycle: 0,
			Mode:  ACC,
			Ins:   ROR,
		}
	case 0x6C:
		return &opcode{
			Cycle: 0,
			Mode:  IND,
			Ins:   JMP,
		}
	case 0x6D:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   ADC,
		}
	case 0x6E:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   ROR,
		}
	case 0x70:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   BVS,
		}
	case 0x71:
		return &opcode{
			Cycle: 0,
			Mode:  IDY,
			Ins:   ADC,
		}
	case 0x75:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   ADC,
		}
	case 0x76:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   ROR,
		}
	case 0x78:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   SEI,
		}
	case 0x79:
		return &opcode{
			Cycle: 0,
			Mode:  ABY,
			Ins:   ADC,
		}
	case 0x7D:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   ADC,
		}
	case 0x7E:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   ROR,
		}
	case 0x81:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   STA,
		}
	case 0x84:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   STY,
		}
	case 0x85:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   STA,
		}
	case 0x86:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   STX,
		}
	case 0x88:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   DEY,
		}
	case 0x8A:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   TXA,
		}
	case 0x8C:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   STY,
		}
	case 0x8D:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   STA,
		}
	case 0x8E:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   STX,
		}
	case 0x90:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   BCC,
		}
	case 0x91:
		return &opcode{
			Cycle: 0,
			Mode:  IDY,
			Ins:   STA,
		}
	case 0x94:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   STY,
		}
	case 0x95:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   STA,
		}
	case 0x96:
		return &opcode{
			Cycle: 0,
			Mode:  ZPY,
			Ins:   STX,
		}
	case 0x98:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   TYA,
		}
	case 0x99:
		return &opcode{
			Cycle: 0,
			Mode:  ABY,
			Ins:   STA,
		}
	case 0x9A:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   TXS,
		}
	case 0x9D:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   STA,
		}
	case 0xA0:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   LDY,
		}
	case 0xA1:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   LDA,
		}
	case 0xA2:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   LDX,
		}
	case 0xA4:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   LDY,
		}
	case 0xA5:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   LDA,
		}
	case 0xA6:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   LDX,
		}
	case 0xA8:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   TAY,
		}
	case 0xA9:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   LDA,
		}
	case 0xAA:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   TAX,
		}
	case 0xAC:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   LDY,
		}
	case 0xAD:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   LDA,
		}
	case 0xAE:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   LDX,
		}
	case 0xB0:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   BCS,
		}
	case 0xB1:
		return &opcode{
			Cycle: 0,
			Mode:  IDY,
			Ins:   LDA,
		}
	case 0xB4:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   LDY,
		}
	case 0xB5:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   LDA,
		}
	case 0xB6:
		return &opcode{
			Cycle: 0,
			Mode:  ZPY,
			Ins:   LDX,
		}
	case 0xB8:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   CLV,
		}
	case 0xB9:
		return &opcode{
			Cycle: 0,
			Mode:  ABY,
			Ins:   LDA,
		}
	case 0xBA:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   TSX,
		}
	case 0xBC:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   LDY,
		}
	case 0xBD:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   LDA,
		}
	case 0xBE:
		return &opcode{
			Cycle: 0,
			Mode:  ABY,
			Ins:   LDX,
		}
	case 0xC0:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   CPY,
		}
	case 0xC1:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   CMP,
		}
	case 0xC4:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   CPY,
		}
	case 0xC5:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   CMP,
		}
	case 0xC6:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   DEC,
		}
	case 0xC8:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   INY,
		}
	case 0xC9:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   CMP,
		}
	case 0xCA:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   DEX,
		}
	case 0xCC:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   CPY,
		}
	case 0xCD:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   CMP,
		}
	case 0xCE:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   DEC,
		}
	case 0xD0:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   BNE,
		}
	case 0xD1:
		return &opcode{
			Cycle: 0,
			Mode:  IDY,
			Ins:   CMP,
		}
	case 0xD5:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   CMP,
		}
	case 0xD6:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   DEC,
		}
	case 0xD8:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   CLD,
		}
	case 0xD9:
		return &opcode{
			Cycle: 0,
			Mode:  ABY,
			Ins:   CMP,
		}
	case 0xDD:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   CMP,
		}
	case 0xDE:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   DEC,
		}
	case 0xE0:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   CPX,
		}
	case 0xE1:
		return &opcode{
			Cycle: 0,
			Mode:  IDX,
			Ins:   SBC,
		}
	case 0xE4:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   CPX,
		}
	case 0xE5:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   SBC,
		}
	case 0xE6:
		return &opcode{
			Cycle: 0,
			Mode:  ZPA,
			Ins:   INC,
		}
	case 0xE8:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   INX,
		}
	case 0xE9:
		return &opcode{
			Cycle: 0,
			Mode:  IMM,
			Ins:   SBC,
		}
	case 0xEA:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   NOP,
		}
	case 0xEC:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   CPX,
		}
	case 0xED:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   SBC,
		}
	case 0xEE:
		return &opcode{
			Cycle: 0,
			Mode:  ABS,
			Ins:   INC,
		}
	case 0xF0:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   BEQ,
		}
	case 0xF1:
		return &opcode{
			Cycle: 0,
			Mode:  IDY,
			Ins:   SBC,
		}
	case 0xF5:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   SBC,
		}
	case 0xF6:
		return &opcode{
			Cycle: 0,
			Mode:  ZPX,
			Ins:   INC,
		}
	case 0xF8:
		return &opcode{
			Cycle: 0,
			Mode:  NIL,
			Ins:   SED,
		}
	case 0xF9:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   SBC,
		}
	case 0xFD:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   SBC,
		}
	case 0xFE:
		return &opcode{
			Cycle: 0,
			Mode:  ABX,
			Ins:   INC,
		}
	}
	panic(fmt.Errorf("%v opcode not found", code))
}
