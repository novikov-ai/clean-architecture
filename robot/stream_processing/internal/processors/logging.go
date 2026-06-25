package processors

import (
	"fmt"

	"clean-architecture/robot/stream_processing/internal/event_store"
	"clean-architecture/robot/stream_processing/internal/events"
)

type LoggingProcessor struct{}

func NewLoggingProcessor(store *event_store.EventStore) *LoggingProcessor {
	p := &LoggingProcessor{}
	store.Subscribe(p.handle)
	return p
}

func (p *LoggingProcessor) handle(evt events.Event) {
	fmt.Printf("[LoggingProcessor] Event logged: %s\n", evt.EventType())
}
