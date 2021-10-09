package cpu6502

func LittleEndianAddr(low, high uint8) uint16 {
	addr := uint16(low)
	addr |= (uint16(high) << 8)
	return addr
}

func IsZero(a uint8) bool {
	return a == 0x00
}

func IsNev(a uint8) bool {
	return a&0x80 != 0
}

func IsPos(a uint8) bool {
	return a&0x80 == 0
}

func IsOverFlow(a, b uint8) bool {
	c := a + b

	if IsPos(a) && IsPos(b) && IsNev(c) {
		return true
	}

	if IsNev(a) && IsNev(b) && IsPos(c) {
		return true
	}

	return false
}

func IsCarry(a, b uint8) bool {
	c := uint16(a) + uint16(b)
	return c > 0x00ff
}
