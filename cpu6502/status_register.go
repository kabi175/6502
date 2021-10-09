package cpu6502

import (
	"errors"
	"fmt"
)

const (
	CARRY     = 0
	ZERO      = iota
	INTERRUPT = iota
	DECIMAL   = iota
	BREAK     = iota
	OVERFLOW  = iota
	SIGN      = iota
)

type FlagRegister interface {
	Set(int, bool)
	Get(int) bool
	Reset()
}

type flagRegister struct {
	carry     bool
	zero      bool
	interrupt bool
	decimal   bool
	brk       bool
	overflow  bool
	sign      bool
}

func NewFlagRegister() FlagRegister {
	return &flagRegister{}
}

func (fr *flagRegister) Set(flag int, state bool) {
	switch flag {
	case CARRY:
		fr.carry = state
	case ZERO:
		fr.zero = state
	case INTERRUPT:
		fr.interrupt = state
	case DECIMAL:
		fr.decimal = state
	case BREAK:
		fr.brk = state
	case OVERFLOW:
		fr.overflow = state
	case SIGN:
		fr.sign = state
	}
	errorMessage := fmt.Sprintf("status register func Set \n %v flag not found", flag)
	panic(errors.New(errorMessage))

}

func (fr *flagRegister) Get(flag int) bool {
	switch flag {
	case CARRY:
		return fr.carry
	case ZERO:
		return fr.zero
	case INTERRUPT:
		return fr.interrupt
	case DECIMAL:
		return fr.decimal
	case BREAK:
		return fr.brk
	case OVERFLOW:
		return fr.overflow
	case SIGN:
		return fr.sign
	}
	errorMessage := fmt.Sprintf("status register func Get \n %v flag not found", flag)
	panic(errors.New(errorMessage))
}

func (fr *flagRegister) Reset() {
	fr.carry = false
	fr.interrupt = false
	fr.decimal = false
	fr.zero = false
	fr.brk = false
	fr.overflow = false
	fr.sign = false
}
