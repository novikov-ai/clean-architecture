package domain

import (
	"fmt"

	"clean-architecture/robot/stream_processing/internal/events"
)

type Command interface {
	ToEvents(robotID string) []events.Event
	CommandType() string
}

type MoveCommand struct{ Distance float64 }

func (c MoveCommand) ToEvents(robotID string) []events.Event {
	return []events.Event{events.MoveRequestedEvent{ID: robotID, Distance: c.Distance}}
}
func (c MoveCommand) CommandType() string { return fmt.Sprintf("MOVE %.2f", c.Distance) }

type TurnCommand struct{ Angle float64 }

func (c TurnCommand) ToEvents(robotID string) []events.Event {
	return []events.Event{events.TurnRequestedEvent{ID: robotID, Angle: c.Angle}}
}
func (c TurnCommand) CommandType() string { return fmt.Sprintf("TURN %.2f", c.Angle) }

type SetStateCommand struct{ Mode CleaningMode }

func (c SetStateCommand) ToEvents(robotID string) []events.Event {
	return []events.Event{events.StateChangeRequestedEvent{ID: robotID, NewState: int(c.Mode)}}
}
func (c SetStateCommand) CommandType() string { return fmt.Sprintf("SET_STATE %d", c.Mode) }

type StartCommand struct{}

func (c StartCommand) ToEvents(robotID string) []events.Event {
	return []events.Event{events.StartRequestedEvent{ID: robotID}}
}
func (c StartCommand) CommandType() string { return "START" }

type StopCommand struct{}

func (c StopCommand) ToEvents(robotID string) []events.Event {
	return []events.Event{events.StopRequestedEvent{ID: robotID}}
}
func (c StopCommand) CommandType() string { return "STOP" }
