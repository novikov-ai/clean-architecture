package processors

import (
	"clean-architecture/robot/stream_processing/internal/domain"
	"clean-architecture/robot/stream_processing/internal/event_store"
	"clean-architecture/robot/stream_processing/internal/events"
	"clean-architecture/robot/stream_processing/internal/projector"
)

type base struct {
	store     *event_store.EventStore
	projector *projector.StateProjector
}

func (b *base) currentState(robotID string) domain.RobotState {
	all := b.store.ForRobot(robotID)
	var stateEvts []events.Event
	for _, e := range all {
		switch e.(type) {
		case events.RobotMovedEvent, events.RobotTurnedEvent, events.RobotStateChangedEvent:
			stateEvts = append(stateEvts, e)
		}
	}
	return b.projector.Project(robotID, stateEvts)
}

func (b *base) emit(evts []events.Event) {
	if len(evts) > 0 {
		b.store.Append(evts)
	}
}
