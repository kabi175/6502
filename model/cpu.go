package model

// An interface for Cpu
type CPU interface {
	Execute(chan bool) // Execute method runs opcodes
	State() []byte     // Get Cpu State in JSON fromat
}

// General Purpose 8-bit Register Interface
type GP8 interface {
	Get() uint8 // Get State
	Set(uint8)  // Set State
}

// A 16-bit Program Counter Interface
//Capable of addressing 64KB's.
type PC16 interface {
	Set(uint16)  // Set PC state
	Get() uint16 // Get PC State
	Increment()  // Increment PC by 0x01
	Decrement()  // Decrement PC by 0x01
	Reset()      // Reset the PC to Initial State
}

// A 8-bit Stack Pointer interface
type SP8 interface {
	Set(uint16)  // Set State
	Get() uint16 // Get State
	Increment()  // Increment Stack by 0x01
	Decrement()  // Decrement Stack by 0x01
	Reset()      // Reset Stack Pointer to Initial state
}

// Flag Register interface
type FlagRegister interface {
	Set(int, bool)
	Get(int) bool
	Byte() uint8
	SetByte(uint8)
	Reset()
}
