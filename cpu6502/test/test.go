package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Cputest struct {
	Prg  []uint8
	PC   uint16
	SP   uint8
	Flag uint8
	X    uint8
	Y    uint8
	A    uint8
}

func ProgramTest(t *testing.T, tests []Cputest) {
	for id, test := range tests {
		testId := fmt.Sprintf("SubTest_%v", id)
		t.Run(testId, func(t *testing.T) {
			cpu := NewCpuBuilder(&Config{Prg: test.Prg})
			cpu.Execute(make(chan bool))
			assert.Equalf(t, test.A, cpu.A.Get(), "Error A Register")
			assert.Equalf(t, test.X, cpu.X.Get(), "Error X Register")
			assert.Equalf(t, test.Y, cpu.Y.Get(), "Error Y Register")
			assert.Equalf(t, test.SP, uint8(cpu.SP.Get()), "Error SP")
			assert.Equalf(t, test.PC, cpu.PC.Get(), "Error PC")
			assert.Equalf(t, test.Flag, cpu.Flag.Byte(), "Error Flag Register:%b", cpu.Flag.Byte())
		})
	}
}
