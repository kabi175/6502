package cpu

type SR uint8

const (
	CARRY SR = 1 << iota
	ZERO
	INTERRUPT
	DECIMAL
	BREAK
	_
	OVERFLOW
	SIGN
)

func (c *CPU6502) SetSR(reg SR, enable bool) {
	if enable {
		c.SR = c.SR | uint8(reg)
	} else {
		c.SR = c.SR & invert(uint8(reg))
	}
}

func (c *CPU6502) GetSR(reg SR) bool {
	return c.SR&uint8(reg) != 0
}
