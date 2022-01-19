package cpu6502

type Event interface {
	Run()
}

type EventQueue struct {
	queue []Event
}

func (e *EventQueue) Run() {
	for _, event := range e.queue {
		event.Run()
	}
}
