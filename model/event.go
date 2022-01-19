package model

// Event interface
// change and store CPU state

type Event interface {
	Run(CPU, Bus16)
}

// Event Queue Interface
// Capable of processing and add events
type EventQueue interface {
	Run(CPU, Bus16)
	AddEvent(events ...Event)
}
