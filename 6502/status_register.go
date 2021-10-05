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

type status struct {
	carry     bool
	zero      bool
	interrupt bool
	decimal   bool
	brk       bool
	overflow  bool
	sign      bool
}

func (this *status) Set(flag int, state bool) {
	switch flag {
	case CARRY:
		this.carry = state
	case ZERO:
		this.zero = state
	case INTERRUPT:
		this.interrupt = state
	case DECIMAL:
		this.decimal = state
	case BREAK:
		this.brk = state
	case OVERFLOW:
		this.overflow = state
	case SIGN:
		this.sign = state
	}
	errorMessage := fmt.Sprintf("status register func Set \n %v flag not found", flag)
	panic(errors.New(errorMessage))

}

func (this *status) Get(flag int) bool {
	switch flag {
	case CARRY:
		return this.carry
	case ZERO:
		return this.zero
	case INTERRUPT:
		return this.interrupt
	case DECIMAL:
		return this.decimal
	case BREAK:
		return this.brk
	case OVERFLOW:
		return this.overflow
	case SIGN:
		return this.sign
	}
	errorMessage := fmt.Sprintf("status register func Get \n %v flag not found", flag)
	panic(errors.New(errorMessage))
}
