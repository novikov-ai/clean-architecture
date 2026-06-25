package projector

import (
	"clean-architecture/robot/stream_processing/internal/domain"
	"clean-architecture/robot/stream_processing/internal/events"
)

type StateProjector struct {
	initial domain.RobotState
}

func New(initial domain.RobotState) *StateProjector {
	return &StateProjector{initial: initial}
}

func (sp *StateProjector) Project(robotID string, evts []events.Event) domain.RobotState {
	state := sp.initial
	for _, evt := range evts {
		if evt.RobotID() != robotID {
			continue
		}
		switch e := evt.(type) {
		case events.RobotMovedEvent:
			state.X = e.NewX
			state.Y = e.NewY
		case events.RobotTurnedEvent:
			state.Angle = e.NewAngle
		case events.RobotStateChangedEvent:
			state.State = e.NewState
		}
	}
	return state
}
