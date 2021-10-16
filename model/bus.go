package model

type Bus16 interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}
