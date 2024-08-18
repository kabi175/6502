package cpu

import "math"

func invert(inp uint8) uint8 {
	return inp ^ 0xff
}

func littleEndian(lo, hi uint8) uint16 {
	return (uint16(hi) << 8) | uint16(lo)
}

func add(op1, op2 uint8) (uint8, bool) {
	sum := uint16(op1) + uint16(op2)
	return uint8(sum), sum > math.MaxUint8
}
