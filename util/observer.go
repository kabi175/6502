package util

import (
	"github.com/kabi175/6502/model"
)

type subject struct {
	callback func(model.State)
}

type Observer struct {
	subjects []*subject
}

func NewObserver() *Observer {
	return &Observer{}
}

func (o *Observer) Subscribe(callback func(model.State)) (func(), error) {
	sub := &subject{callback}
	o.subjects = append(o.subjects, sub)
	return func() {
		for i, iter := range o.subjects {
			if sub == iter {
				o.subjects = append(o.subjects[:i], o.subjects[i+1:]...)
				break
			}
		}
	}, nil
}

func (o *Observer) Publish(state model.State) error {
	for _, sub := range o.subjects {
		sub.callback(state)
	}
	return nil
}
