package events

import "fmt"

type Event interface {
	EventType() string
	RobotID() string
}

type MoveRequestedEvent struct {
	ID       string
	Distance float64
}

func (e MoveRequestedEvent) EventType() string { return fmt.Sprintf("MOVE_REQUESTED %.2f", e.Distance) }
func (e MoveRequestedEvent) RobotID() string   { return e.ID }

type TurnRequestedEvent struct {
	ID    string
	Angle float64
}

func (e TurnRequestedEvent) EventType() string { return fmt.Sprintf("TURN_REQUESTED %.2f", e.Angle) }
func (e TurnRequestedEvent) RobotID() string   { return e.ID }

type StateChangeRequestedEvent struct {
	ID       string
	NewState int
}

func (e StateChangeRequestedEvent) EventType() string {
	return fmt.Sprintf("STATE_CHANGE_REQUESTED %d", e.NewState)
}
func (e StateChangeRequestedEvent) RobotID() string { return e.ID }

type StartRequestedEvent struct{ ID string }

func (e StartRequestedEvent) EventType() string { return "START_REQUESTED" }
func (e StartRequestedEvent) RobotID() string   { return e.ID }

type StopRequestedEvent struct{ ID string }

func (e StopRequestedEvent) EventType() string { return "STOP_REQUESTED" }
func (e StopRequestedEvent) RobotID() string   { return e.ID }

type RobotMovedEvent struct {
	ID                          string
	OldX, OldY, NewX, NewY, Distance float64
}

func (e RobotMovedEvent) EventType() string {
	return fmt.Sprintf("ROBOT_MOVED from (%.2f, %.2f) to (%.2f, %.2f)", e.OldX, e.OldY, e.NewX, e.NewY)
}
func (e RobotMovedEvent) RobotID() string { return e.ID }

type RobotTurnedEvent struct {
	ID                 string
	OldAngle, NewAngle float64
}

func (e RobotTurnedEvent) EventType() string {
	return fmt.Sprintf("ROBOT_TURNED from %.2f to %.2f", e.OldAngle, e.NewAngle)
}
func (e RobotTurnedEvent) RobotID() string { return e.ID }

type RobotStateChangedEvent struct {
	ID                 string
	OldState, NewState int
}

func (e RobotStateChangedEvent) EventType() string {
	return fmt.Sprintf("ROBOT_STATE_CHANGED from %d to %d", e.OldState, e.NewState)
}
func (e RobotStateChangedEvent) RobotID() string { return e.ID }

type RobotStartedEvent struct{ ID string }

func (e RobotStartedEvent) EventType() string { return "ROBOT_STARTED" }
func (e RobotStartedEvent) RobotID() string   { return e.ID }

type RobotStoppedEvent struct{ ID string }

func (e RobotStoppedEvent) EventType() string { return "ROBOT_STOPPED" }
func (e RobotStoppedEvent) RobotID() string   { return e.ID }
