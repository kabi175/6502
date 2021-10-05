package cpu6502

type Instruction struct {
	Cycle uint8
	Mode  func(*cpu6502) uint8
	Ins   func(*cpu6502) uint8
}

func load() map[uint16]Instruction {
	opCodes := make(map[uint16]Instruction)
	return opCodes
}
