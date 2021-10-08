package cpu6502

type Instruction struct {
	Cycle uint8
	Mode  func(*cpu6502) uint8
	Ins   func(*cpu6502) uint8
}

func (i *Instruction) Execute(c *cpu6502) uint8 {
	cycles := i.Cycle
	cycles += i.Mode(c)
	cycles += i.Ins(c)
	return cycles
}

func loadInstructionSet() map[uint16]Instruction {
	opCodes := make(map[uint16]Instruction)
	return opCodes
}

func ADC(c *cpu6502) uint8 {
	return 0
}

func AND(c *cpu6502) uint8 {
	return 0
}

func ASL(c *cpu6502) uint8 {
	return 0
}

func BCC(c *cpu6502) uint8 {
	return 0
}

func BCS(c *cpu6502) uint8 {
	return 0
}
func BEQ(c *cpu6502) uint8 {
	return 0
}

func BIT(c *cpu6502) uint8 {
	return 0
}

func BMI(c *cpu6502) uint8 {
	return 0
}
func BNE(c *cpu6502) uint8 {
	return 0
}
func BPL(c *cpu6502) uint8 {
	return 0
}

func BRK(c *cpu6502) uint8 {
	return 0
}

func BVC(c *cpu6502) uint8 {
	return 0
}
func BVS(c *cpu6502) uint8 {
	return 0
}
func CLC(c *cpu6502) uint8 {
	return 0
}
func CLD(c *cpu6502) uint8 {
	return 0
}
func CLI(c *cpu6502) uint8 {
	return 0
}
func CLV(c *cpu6502) uint8 {
	return 0
}
func CMP(c *cpu6502) uint8 {
	return 0
}
func CPX(c *cpu6502) uint8 {
	return 0
}
func DEC(c *cpu6502) uint8 {
	return 0
}
func DEX(c *cpu6502) uint8 {
	return 0
}
func DEY(c *cpu6502) uint8 {
	return 0
}
func EOR(c *cpu6502) uint8 {
	return 0
}
func INC(c *cpu6502) uint8 {
	return 0
}
func INX(c *cpu6502) uint8 {
	return 0
}
func INY(c *cpu6502) uint8 {
	return 0
}
func JMP(c *cpu6502) uint8 {
	return 0
}
func JSR(c *cpu6502) uint8 {
	return 0
}
func LDA(c *cpu6502) uint8 {
	return 0
}
func LDX(c *cpu6502) uint8 {
	return 0
}
func LDY(c *cpu6502) uint8 {
	return 0
}
func LSR(c *cpu6502) uint8 {
	return 0
}
func NOP(c *cpu6502) uint8 {
	return 0
}
func ORA(c *cpu6502) uint8 {
	return 0
}
func PHA(c *cpu6502) uint8 {
	return 0
}
func PHP(c *cpu6502) uint8 {
	return 0
}
func PLA(c *cpu6502) uint8 {
	return 0
}
func PLP(c *cpu6502) uint8 {
	return 0
}
func ROL(c *cpu6502) uint8 {
	return 0
}
func ROR(c *cpu6502) uint8 {
	return 0
}
func RTI(c *cpu6502) uint8 {
	return 0
}
func RTS(c *cpu6502) uint8 {
	return 0
}
func SBC(c *cpu6502) uint8 {
	return 0
}
func SEC(c *cpu6502) uint8 {
	return 0
}
func SED(c *cpu6502) uint8 {
	return 0
}
func SEI(c *cpu6502) uint8 {
	return 0
}
func STA(c *cpu6502) uint8 {
	return 0
}
func STX(c *cpu6502) uint8 {
	return 0
}
func STY(c *cpu6502) uint8 {
	return 0
}
func TAX(c *cpu6502) uint8 {
	return 0
}
func TAY(c *cpu6502) uint8 {
	return 0
}
func TSX(c *cpu6502) uint8 {
	return 0
}
func TXA(c *cpu6502) uint8 {
	return 0
}
func TXS(c *cpu6502) uint8 {
	return 0
}
func TYA(c *cpu6502) uint8 {
	return 0
}
