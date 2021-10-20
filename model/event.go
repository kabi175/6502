package model

// Event interface
// change and store CPU state
type Event interface {
	Run()
}

// Event Queue Interface
// Capable of processing and add events
type EventQueue interface {
	Push(Event)
	Consume()
}
