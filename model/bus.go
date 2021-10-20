package model

// An interface for 16-bit BUS
type Bus16 interface {
	Read(uint16) uint8   // Read  8-bitData from RAM
	Write(uint16, uint8) // Write 8-bit Data to Ram
}
