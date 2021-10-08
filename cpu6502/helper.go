package cpu6502

func littleEndianAddr(low, high uint8) uint16 {
	addr := uint16(low)
	addr |= (uint16(high) << 8)
	return addr
}
