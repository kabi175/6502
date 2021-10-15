package util

type Subject interface {
	OnStateChange()
}

type Observer interface {
	Subscribe(Subject)
	UnSubscribe(Subject)
}

type ObserverImp struct {
	subjects []Subject
}

func (o *ObserverImp) Subscribe(s Subject) {
	o.subjects = append(o.subjects, s)
}

func (o *ObserverImp) UnSubscribe(s Subject) {
	for index, sub := range o.subjects {
		if sub == s {
			o.subjects = remove(o.subjects, index)
			return
		}
	}
}

func remove(s []Subject, index int) []Subject {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}
