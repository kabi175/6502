package bus

type Bus16 struct {
	ram [1024 * 64]uint8
}

func NewBus16(prg []uint8) *Bus16 {
	var ram [1024 * 64]uint8
	for index, code := range prg {
		ram[index] = code
	}
	return &Bus16{
		ram: ram,
	}
}

func (b *Bus16) Read(add uint16) uint8 {
	return b.ram[add]
}

func (b *Bus16) Write(add uint16, data uint8) {
	b.ram[add] = data
}
