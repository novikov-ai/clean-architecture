package processors

import (
	"fmt"
	"math"

	"clean-architecture/robot/stream_processing/internal/event_store"
	"clean-architecture/robot/stream_processing/internal/events"
	"clean-architecture/robot/stream_processing/internal/projector"
)

type MovementProcessor struct {
	base
}

func NewMovementProcessor(store *event_store.EventStore, proj *projector.StateProjector) *MovementProcessor {
	p := &MovementProcessor{base: base{store: store, projector: proj}}
	store.Subscribe(p.handle)
	return p
}

func (p *MovementProcessor) handle(evt events.Event) {
	switch e := evt.(type) {
	case events.MoveRequestedEvent:
		p.handleMove(e)
	case events.TurnRequestedEvent:
		p.handleTurn(e)
	}
}

func (p *MovementProcessor) handleMove(evt events.MoveRequestedEvent) {
	fmt.Printf("[MovementProcessor] Processing move request for robot %s\n", evt.ID)

	state := p.currentState(evt.ID)
	angleRads := state.Angle * (math.Pi / 180.0)
	newX := state.X + evt.Distance*math.Cos(angleRads)
	newY := state.Y + evt.Distance*math.Sin(angleRads)

	p.emit([]events.Event{events.RobotMovedEvent{
		ID:       evt.ID,
		OldX:     state.X,
		OldY:     state.Y,
		NewX:     newX,
		NewY:     newY,
		Distance: evt.Distance,
	}})
}

func (p *MovementProcessor) handleTurn(evt events.TurnRequestedEvent) {
	fmt.Printf("[MovementProcessor] Processing turn request for robot %s\n", evt.ID)

	state := p.currentState(evt.ID)

	p.emit([]events.Event{events.RobotTurnedEvent{
		ID:       evt.ID,
		OldAngle: state.Angle,
		NewAngle: state.Angle + evt.Angle,
	}})
}
