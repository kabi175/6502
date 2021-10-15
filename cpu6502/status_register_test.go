package cpu6502

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusRegister(t *testing.T) {
	sr := NewFlagRegister()
	sr.Reset()
	tests := []struct {
		arg uint8
	}{
		{arg: 0b00100000},
		{arg: 0b10101000},
		{arg: 0b10101010},
		{arg: 0b11101010},
		{arg: 0b10100101},
		{arg: 0b11111111},
		{arg: 0b10100100},
		{arg: 0b10101111},
		{arg: 0b11110000},
	}
	for id, test := range tests {
		test_id := fmt.Sprintf("Test Id:%v", id)
		t.Run(test_id, func(t *testing.T) {
			sr.SetByte(test.arg)
			got := sr.Byte()
			assert.Equal(t, test.arg, got)

		})
	}
}
