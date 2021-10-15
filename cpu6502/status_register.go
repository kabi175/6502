package cpu6502

import (
	"errors"
	"fmt"
)

const (
	CARRY = iota
	ZERO
	INTERRUPT
	DECIMAL
	BREAK
	OVERFLOW
	SIGN
)

type flagRegister struct {
	carry     bool
	zero      bool
	interrupt bool
	decimal   bool
	brk       bool
	unused    bool
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
	default:
		errorMessage := fmt.Sprintf("status register func Set \n %v flag not found", flag)
		panic(errors.New(errorMessage))
	}
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
	default:
		errorMessage := fmt.Sprintf("status register func Get \n %v flag not found", flag)
		panic(errors.New(errorMessage))
	}
}

func (fr *flagRegister) Reset() {
	fr.carry = false
	fr.interrupt = false
	fr.decimal = false
	fr.zero = false
	fr.brk = false
	fr.unused = true
	fr.overflow = false
	fr.sign = false
}

func (fr *flagRegister) Byte() uint8 {
	flag := uint8(0x20)
	if fr.carry {
		flag |= 0x01
	}
	if fr.zero {
		flag |= 0x02
	}
	if fr.interrupt {
		flag |= 0x04
	}
	if fr.decimal {
		flag |= 0x08
	}
	if fr.brk {
		flag |= 0x10
	}
	if fr.overflow {
		flag |= 0x40
	}
	if fr.sign {
		flag |= 0x80
	}
	return flag
}

func (fr *flagRegister) SetByte(flag uint8) {

	fr.Reset()

	if flag&0x01 != 0 {
		fr.carry = true
	}
	if flag&0x02 != 0 {
		fr.zero = true
	}
	if flag&0x04 != 0 {
		fr.interrupt = true
	}
	if flag&0x08 != 0 {
		fr.decimal = true
	}
	if flag&0x10 != 0 {
		fr.brk = true
	}
	if flag&0x40 != 0 {
		fr.overflow = true
	}
	if flag&0x80 != 0 {
		fr.sign = true
	}
}
