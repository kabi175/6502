package cpu6502

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kabi175/6502/bus"
)

func TestStackPointer(t *testing.T) {
	bus := bus.NewBus16([]uint8{})
	cpu := newCPU(bus)
	t.Run("InitialValue", func(t *testing.T) {
		assert.Equal(t, uint16(0x00FF), cpu.SP.Get())
	})
	t.Run("PUSH", func(t *testing.T) {
		cpu.push(uint8(0xFF))
		assert.Equal(t, uint16(0x00FF-1), cpu.SP.Get())
	})
	t.Run("PULL", func(t *testing.T) {
		byte := cpu.pull()
		assert.Equal(t, uint16(0x00FF), cpu.SP.Get())
		assert.Equal(t, uint8(0xFF), byte)
	})
}
