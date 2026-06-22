package event_store

import "clean-architecture/event_sourcing/internal/domain"

type EventStore struct {
	log []Event
}

func New(events ...Event) EventStore {
	return EventStore{
		log: events,
	}
}

func (es EventStore) Latest() Event {
	if len(es.log) == 0 {
		return Event{}
	}

	return es.log[len(es.log)-1]
}

func (es EventStore) Undo() Event {
	if len(es.log) == 0 {
		return Event{}
	}

	event := es.log[len(es.log)-1]

	es.log = es.log[:len(es.log)-1]

	return event
}

func (es EventStore) Redo(event Event) {
	es.log = append(es.log, event)
}

func (es EventStore) Produce(event Event) {
	es.log = append(es.log, event)
}

func (es EventStore) Status() domain.Robot{
	robot := domain.Robot{}
	for _, event := range es.log{
		robot = event.Command.Execute(robot)
	}

	return robot
}