package model

type Event interface {
	Run()
}

type EventQueue interface {
	Push(Event)
	Consume()
}
