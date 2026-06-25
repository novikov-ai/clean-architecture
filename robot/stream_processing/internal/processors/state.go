package processors

import (
	"fmt"

	"clean-architecture/robot/stream_processing/internal/event_store"
	"clean-architecture/robot/stream_processing/internal/events"
	"clean-architecture/robot/stream_processing/internal/projector"
)

type StateProcessor struct {
	base
}

func NewStateProcessor(store *event_store.EventStore, proj *projector.StateProjector) *StateProcessor {
	p := &StateProcessor{base: base{store: store, projector: proj}}
	store.Subscribe(p.handle)
	return p
}

func (p *StateProcessor) handle(evt events.Event) {
	switch e := evt.(type) {
	case events.StateChangeRequestedEvent:
		p.handleStateChange(e)
	case events.StartRequestedEvent:
		p.handleStart(e)
	case events.StopRequestedEvent:
		p.handleStop(e)
	}
}

func (p *StateProcessor) handleStateChange(evt events.StateChangeRequestedEvent) {
	fmt.Printf("[StateProcessor] Processing state change request for robot %s\n", evt.ID)

	state := p.currentState(evt.ID)

	p.emit([]events.Event{events.RobotStateChangedEvent{
		ID:       evt.ID,
		OldState: state.State,
		NewState: evt.NewState,
	}})
}

func (p *StateProcessor) handleStart(evt events.StartRequestedEvent) {
	fmt.Printf("[StateProcessor] Processing start request for robot %s\n", evt.ID)
	p.emit([]events.Event{events.RobotStartedEvent{ID: evt.ID}})
}

func (p *StateProcessor) handleStop(evt events.StopRequestedEvent) {
	fmt.Printf("[StateProcessor] Processing stop request for robot %s\n", evt.ID)
	p.emit([]events.Event{events.RobotStoppedEvent{ID: evt.ID}})
}
