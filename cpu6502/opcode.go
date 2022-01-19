package cpu6502

type ADDRMODE string
type INSTRUCTION string

// ADDRESSING MODE ENUM
const (
	ACC ADDRMODE = "ACC"
	IMM ADDRMODE = "IMM"
	ZPA ADDRMODE = "ZPA"
	ZPX ADDRMODE = "ZPX"
	ZPY ADDRMODE = "ZPY"
	ABS ADDRMODE = "ABS"
	ABX ADDRMODE = "ABX"
	ABY ADDRMODE = "ABY"
	IMP ADDRMODE = "IMP"
	REL ADDRMODE = "REL"
	IDX ADDRMODE = "IDX"
	IDY ADDRMODE = "IDY"
	IND ADDRMODE = "IND"
)

// INSTRUCTION ENUM
const (
	ADC INSTRUCTION = "ADC"
	AND INSTRUCTION = "AND"
	ASL INSTRUCTION = "ASL"
	BCC INSTRUCTION = "BCC"
	BCS INSTRUCTION = "BCS"
	BEQ INSTRUCTION = "BEQ"
	BIT INSTRUCTION = "BIT"
	BMI INSTRUCTION = "BMI"
	BNE INSTRUCTION = "BNE"
	BPL INSTRUCTION = "BPL"
	BRK INSTRUCTION = "BRK"
	BVC INSTRUCTION = "BVC"
	BVS INSTRUCTION = "BVS"
	CLC INSTRUCTION = "CLC"
	CLD INSTRUCTION = "CLD"
	CLI INSTRUCTION = "CLI"
	CLV INSTRUCTION = "CLV"
	CMP INSTRUCTION = "CMP"
	CPX INSTRUCTION = "CPX"
	CPY INSTRUCTION = "CPY"
	DEC INSTRUCTION = "DEC"
	DEX INSTRUCTION = "DEX"
	DEY INSTRUCTION = "DEY"
	EOR INSTRUCTION = "EOR"
	INC INSTRUCTION = "INC"
	INX INSTRUCTION = "INX"
	INY INSTRUCTION = "INY"
	JMP INSTRUCTION = "JMP"
	JSR INSTRUCTION = "JSR"
	LDA INSTRUCTION = "LDA"
	LDX INSTRUCTION = "LDX"
	LDY INSTRUCTION = "LDY"
	LSR INSTRUCTION = "LSR"
	NOP INSTRUCTION = "NOP"
	ORA INSTRUCTION = "ORA"
	PHA INSTRUCTION = "PHA"
	PHP INSTRUCTION = "PHP"
	PLA INSTRUCTION = "PLA"
	PLP INSTRUCTION = "PLP"
	ROL INSTRUCTION = "ROL"
	ROR INSTRUCTION = "ROR"
	RTI INSTRUCTION = "RTI"
	RTS INSTRUCTION = "RTS"
	SBC INSTRUCTION = "SBC"
	SEC INSTRUCTION = "SEC"
	SED INSTRUCTION = "SED"
	SEI INSTRUCTION = "SEI"
	STA INSTRUCTION = "STA"
	STX INSTRUCTION = "STX"
	STY INSTRUCTION = "STY"
	TAX INSTRUCTION = "TAX"
	TAY INSTRUCTION = "TAY"
	TSX INSTRUCTION = "TSX"
	TXA INSTRUCTION = "TXA"
	TXS INSTRUCTION = "TXS"
	TYA INSTRUCTION = "TYA"
)

type opcode struct {
	Cycle uint8
	Mode  ADDRMODE
	Ins   INSTRUCTION
}

func (o *opcode) Execute(cpu interface{}) uint8 {
	c, ok := cpu.(*CPU6502)
	if !ok {
		panic("Invalid CPU type")
	}
	cycle := o.Cycle
	switch o.Mode {
	case ACC:
		cycle += o.ACC(c)
	case IMM:
		cycle += o.IMM(c)
	case ZPA:
		cycle += o.ZPA(c)
	case ZPX:
		cycle += o.ZPX(c)
	case ZPY:
		cycle += o.ZPY(c)
	case ABS:
		cycle += o.ABS(c)
	case ABX:
		cycle += o.ABX(c)
	case ABY:
		cycle += o.ABY(c)
	case IMP:
		cycle += o.IMP(c)
	case REL:
		cycle += o.REL(c)
	case IDX:
		cycle += o.IDX(c)
	case IDY:
		cycle += o.IDY(c)
	case IND:
		cycle += o.IND(c)
	}
	switch o.Ins {
	case ADC:
		cycle += o.ADC(c)
	case AND:
		cycle += o.AND(c)
	case ASL:
		cycle += o.ASL(c)
	case BCC:
		cycle += o.BCC(c)
	case BCS:
		cycle += o.BCS(c)
	case BEQ:
		cycle += o.BEQ(c)
	case BIT:
		cycle += o.BIT(c)
	case BMI:
		cycle += o.BMI(c)
	case BNE:
		cycle += o.BNE(c)
	case BPL:
		cycle += o.BPL(c)
	case BRK:
		cycle += o.BRK(c)
	case BVC:
		cycle += o.BVC(c)
	case BVS:
		cycle += o.BVS(c)
	case CLC:
		cycle += o.CLC(c)
	case CLD:
		cycle += o.CLD(c)
	case CLI:
		cycle += o.CLI(c)
	case CLV:
		cycle += o.CLV(c)
	case CMP:
		cycle += o.CMP(c)
	case CPX:
		cycle += o.CPX(c)
	case CPY:
		cycle += o.CPY(c)
	case DEC:
		cycle += o.DEC(c)
	case DEX:
		cycle += o.DEX(c)
	case DEY:
		cycle += o.DEY(c)
	case EOR:
		cycle += o.EOR(c)
	case INC:
		cycle += o.INC(c)
	case INX:
		cycle += o.INX(c)
	case INY:
		cycle += o.INY(c)
	case JMP:
		cycle += o.JMP(c)
	case JSR:
		cycle += o.JSR(c)
	case LDA:
		cycle += o.LDA(c)
	case LDX:
		cycle += o.LDX(c)
	case LDY:
		cycle += o.LDY(c)
	case LSR:
		cycle += o.LSR(c)
	case NOP:
		cycle += o.NOP(c)
	case ORA:
		cycle += o.ORA(c)
	case PHA:
		cycle += o.PHA(c)
	case PHP:
		cycle += o.PHP(c)
	case PLA:
		cycle += o.PLA(c)
	case PLP:
		cycle += o.PLP(c)
	case ROL:
		cycle += o.ROL(c)
	case ROR:
		cycle += o.ROR(c)
	case RTI:
		cycle += o.RTI(c)
	case RTS:
		cycle += o.RTS(c)
	case SBC:
		cycle += o.SBC(c)
	case SEC:
		cycle += o.SEC(c)
	case SED:
		cycle += o.SED(c)
	case SEI:
		cycle += o.SEI(c)
	case STA:
		cycle += o.STA(c)
	case STX:
		cycle += o.STX(c)
	case STY:
		cycle += o.STY(c)
	case TAX:
		cycle += o.TAX(c)
	case TAY:
		cycle += o.TAY(c)
	case TSX:
		cycle += o.TSX(c)
	case TXA:
		cycle += o.TXA(c)
	case TXS:
		cycle += o.TXS(c)
	case TYA:
		cycle += o.TYA(c)
	}
	return cycle
}

func (o *opcode) IsBreak() bool {
	return o.Ins == BRK
}

func (o *opcode) State() []byte {
	panic("Not implemented")
}
