package cpu6502

import "github.com/kabi175/6502/model"

type EventQueue struct {
	queue []model.Event
}

func (e *EventQueue) Run(cpu model.CPU, bus model.Bus16) {
	for _, event := range e.queue {
		event.Run(cpu, bus)
	}
}

func (e *EventQueue) AddEvent(event ...model.Event) {
	e.queue = append(e.queue, event...)
}
