package bus

type Bus16 struct {
	Ram [1024 * 64]uint8
}

func NewBus16(prg []uint8) *Bus16 {
	var ram [1024 * 64]uint8
	for index, code := range prg {
		ram[index] = code
	}
	return &Bus16{
		Ram: ram,
	}
}

func (b *Bus16) Read(add uint16) uint8 {
	if add > 0x0000 && add < 0xFFFF {
		return b.Ram[add]
	}
	return 0x00
}

func (b *Bus16) Write(add uint16, data uint8) {
	if add > 0x0000 && add < 0xFFFF {
		b.Ram[add] = data
	}
}
