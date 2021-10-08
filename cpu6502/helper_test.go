package cpu6502

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLittleEndianAddr(t *testing.T) {
	tests := []struct {
		low  uint8
		high uint8
		addr uint16
	}{
		{low: 0x00, high: 0x00, addr: 0x0000},
		{low: 0x40, high: 0x32, addr: 0x3240},
		{low: 0x43, high: 0x42, addr: 0x4243},
		{low: 0xff, high: 0xff, addr: 0xffff},
		{low: 0xad, high: 0xdf, addr: 0xdfad},
		{low: 0xcd, high: 0xdc, addr: 0xdccd},
		{low: 0x5f, high: 0x00, addr: 0x005f},
	}
	for _, test := range tests {
		addr := littleEndianAddr(test.low, test.high)
		assert.Equal(t, addr, test.addr)
	}
}
