package event_store

import "clean-architecture/event_sourcing/internal/domain"

type Event struct {
	Command domain.Commander
}

func NewEvent(cmd domain.Commander) Event {
	return Event{
		Command: cmd,
	}
}
