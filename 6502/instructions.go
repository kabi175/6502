package cpu6502

type Instruction struct {
	Cycle uint8
	Mode  func() uint8
	Ins   func() uint8
}

func load() map[uint8]Instruction {
	opCodes := make(map[uint8]Instruction)
	return opCodes
}
