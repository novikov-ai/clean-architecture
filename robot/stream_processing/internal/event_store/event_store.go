package event_store

import (
	"sync"

	"clean-architecture/robot/stream_processing/internal/events"
)

type Subscriber func(events.Event)

type EventStore struct {
	mu          sync.Mutex
	all         []events.Event
	subscribers []Subscriber
}

func New() *EventStore {
	return &EventStore{}
}

func (es *EventStore) Append(evts []events.Event) {
	es.mu.Lock()
	es.all = append(es.all, evts...)
	subs := make([]Subscriber, len(es.subscribers))
	copy(subs, es.subscribers)
	es.mu.Unlock()

	for _, evt := range evts {
		for _, sub := range subs {
			sub(evt)
		}
	}
}

func (es *EventStore) All() []events.Event {
	es.mu.Lock()
	defer es.mu.Unlock()
	result := make([]events.Event, len(es.all))
	copy(result, es.all)
	return result
}

func (es *EventStore) ForRobot(robotID string) []events.Event {
	es.mu.Lock()
	defer es.mu.Unlock()
	var result []events.Event
	for _, e := range es.all {
		if e.RobotID() == robotID {
			result = append(result, e)
		}
	}
	return result
}

func (es *EventStore) Subscribe(sub Subscriber) {
	es.mu.Lock()
	defer es.mu.Unlock()
	es.subscribers = append(es.subscribers, sub)
}
