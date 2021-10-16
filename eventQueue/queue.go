package eventqueue

import (
	"sync"

	"github.com/kabi175/6502/model"
	"github.com/oleiade/lane"
)

func NewEventQueue() model.EventQueue {
	return &EventQueue{
		queue: lane.NewQueue(),
	}
}

type EventQueue struct {
	queue *lane.Queue
	mutex sync.Mutex
}

func (q *EventQueue) Push(event model.Event) {
	q.mutex.Lock()
	q.queue.Enqueue(event)
	q.mutex.Unlock()
}

func (q *EventQueue) Pop() model.Event {
	q.mutex.Lock()
	event := q.queue.Dequeue()
	q.mutex.Unlock()
	e, ok := event.(model.Event)
	if !ok {
		panic("Queue Element is not of type Event")
	}
	return e
}

func (q *EventQueue) IsEmpty() bool {
	return q.queue.Empty()
}

func (q *EventQueue) Consume() {
	q.mutex.Lock()
	for !q.IsEmpty() {
		event := q.queue.Dequeue()
		e, ok := event.(model.Event)
		if !ok {
			panic("Queue Element is not of type Event")
		}
		e.Run()
	}
	q.mutex.Unlock()
}
