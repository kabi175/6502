package cpu6502

import (
	"fmt"
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
	for id, test := range tests {
		testID := fmt.Sprintf("subtest-%v", id)
		t.Run(testID, func(t *testing.T) {
			got := LittleEndianAddr(test.low, test.high)
			assert.Equal(t, test.addr, got)
		})
	}
}

func TestIsOverFlow(t *testing.T) {
	tests := []struct {
		a   uint8
		b   uint8
		exp bool
	}{
		{a: 0x50, b: 0x10, exp: false},
		{a: 0x50, b: 0x50, exp: true},
		{a: 0x50, b: 0x90, exp: false},
		{a: 0x50, b: 0xd0, exp: false},
		{a: 0xd0, b: 0x10, exp: false},
		{a: 0xd0, b: 0x50, exp: false},
		{a: 0xd0, b: 0x90, exp: true},
		{a: 0xd0, b: 0xd0, exp: false},
	}
	for id, test := range tests {
		testID := fmt.Sprintf("subtest-%v", id)
		t.Run(testID, func(t *testing.T) {
			got := IsOverFlow(test.a, test.b)
			t.Logf("%X + %X -> %X", test.a, test.b, test.a+test.b)
			assert.Equal(t, test.exp, got)
		})
	}
}

func TestIsPos(t *testing.T) {
	tests := []struct {
		a   uint8
		exp bool
	}{
		{a: 0xff, exp: false},
		{a: 0x0f, exp: true},
		{a: 0xf0, exp: false},
	}
	for _, test := range tests {
		got := IsPos(test.a)
		assert.Equal(t, test.exp, got)
	}
}

func TestIsNev(t *testing.T) {
	tests := []struct {
		a   uint8
		exp bool
	}{
		{a: 0xff, exp: true},
		{a: 0x0f, exp: false},
		{a: 0xf0, exp: true},
	}
	for _, test := range tests {
		got := IsNev(test.a)
		assert.Equal(t, test.exp, got)
	}
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		a   uint8
		exp bool
	}{
		{a: 0xff, exp: false},
		{a: 0x0f, exp: false},
		{a: 0xf0, exp: false},
		{a: 0x00, exp: true},
		{a: 0, exp: true},
	}
	for _, test := range tests {
		got := IsZero(test.a)
		assert.Equal(t, test.exp, got)
	}
}

func TestIsCarry(t *testing.T) {
	tests := []struct {
		a   uint8
		b   uint8
		exp bool
	}{
		{a: 0x50, b: 0x10, exp: false},
		{a: 0x50, b: 0x50, exp: false},
		{a: 0x50, b: 0x90, exp: false},
		{a: 0x50, b: 0xd0, exp: true},
		{a: 0xd0, b: 0x10, exp: false},
		{a: 0xd0, b: 0x50, exp: true},
		{a: 0xd0, b: 0x90, exp: true},
		{a: 0xd0, b: 0xd0, exp: true},
	}
	for id, test := range tests {
		testID := fmt.Sprintf("subtest-%v", id)
		t.Run(testID, func(t *testing.T) {
			got := IsCarry(test.a, test.b)
			t.Logf("%X + %X -> %X", test.a, test.b, test.a+test.b)
			assert.Equal(t, test.exp, got)
		})
	}

}
