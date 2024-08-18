package cpu

type AddressingMode func(cpu *CPU6502) (*uint8, *uint16, uint8)

// Accumulator
// instructions act directly on Accumulator
func ACC(cpu *CPU6502) (opr *uint8, adrr *uint16, cycle uint8) {
	opr = &cpu.AC
	return
}

// Implied Addressing
// instructions act directly on one or more registers or flags internal to the CPU
func IMP(cpu *CPU6502) (opr *uint8, adrr *uint16, cycle uint8) {
	return
}

// Immediate Addressing
// operand is given immediately after the instruction
func IMM(cpu *CPU6502) (*uint8, *uint16, uint8) {
	opr := cpu.FetchData()
	return &opr, nil, 1
}

// Absolute Addressing
func ABS(cpu *CPU6502) (*uint8, *uint16, uint8) {
	lo := cpu.FetchData()
	hi := cpu.FetchData()

	addr := littleEndian(lo, hi)
	opr := cpu.Read(addr)

	return &opr, &addr, 3
}

// Zero-Page Addressing
func ZP(cpu *CPU6502) (*uint8, *uint16, uint8) {
	lo := cpu.FetchData()
	hi := uint8(0)

	addr := littleEndian(lo, hi)
	opr := cpu.Read(addr)

	return &opr, &addr, 2
}

// Absolute Addressing - X
func ABS_X(cpu *CPU6502) (*uint8, *uint16, uint8) {
	lo := cpu.FetchData()
	hi := cpu.FetchData()
	cycle := uint8(3)

	lo, carry := add(lo, cpu.X)
	if carry {
		hi++
		cycle++
	}
	addr := littleEndian(lo, hi)

	opr := cpu.Read(addr)

	return &opr, &addr, cycle
}

// Absolute Addressing - Y
func ABS_Y(cpu *CPU6502) (*uint8, *uint16, uint8) {
	lo := cpu.FetchData()
	hi := cpu.FetchData()
	cycle := uint8(3)

	lo, carry := add(lo, cpu.Y)
	if carry {
		hi++
		cycle++
	}
	addr := littleEndian(lo, hi)

	opr := cpu.Read(addr)

	return &opr, &addr, cycle
}

// Zero-Page Addressing X
func ZP_X(cpu *CPU6502) (*uint8, *uint16, uint8) {
	lo := cpu.FetchData()
	hi := uint8(0)

	lo, _ = add(lo, cpu.X)
	addr := littleEndian(lo, hi)

	opr := cpu.Read(addr)

	return &opr, &addr, 2
}

// Zero-Page Addressing Y
func ZP_Y(cpu *CPU6502) (*uint8, *uint16, uint8) {
	lo := cpu.FetchData()
	hi := uint8(0)

	lo, _ = add(lo, cpu.Y)
	addr := littleEndian(lo, hi)

	opr := cpu.Read(addr)

	return &opr, &addr, 2
}

// Relative Addressing
func REL(cpu *CPU6502) (*uint8, *uint16, uint8) {
	opr := cpu.FetchData()
	return &opr, nil, 1
}

// Indirect Addressing
func IND(cpu *CPU6502) (opr *uint8, adrr *uint16, cycle uint8) {
	lo := cpu.FetchData()
	hi := cpu.FetchData()

	addrLow := littleEndian(lo, hi)
	lo = cpu.Read(addrLow)
	hi = cpu.Read(addrLow + 1)

	addr := littleEndian(lo, hi)
	return nil, &addr, 4
}
