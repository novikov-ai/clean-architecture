package robo

import (
	"clean-architecture/event_sourcing/internal/domain"
	"fmt"
)

type MoveCommand struct {
	value int
}

func NewMoveCommand(value int) MoveCommand {
	return MoveCommand{
		value: value,
	}
}

func (mc MoveCommand) Execute(r domain.Robot) domain.Robot {
	angle := ((r.Angle % 360) + 360) % 360

	robot := r

	switch angle {
	case 0:
		robot.Position.X += mc.value
	case 90:
		robot.Position.Y += mc.value
	case 180:
		robot.Position.X -= mc.value
	case 270:
		robot.Position.Y -= mc.value
	}

	fmt.Printf("POS %v %v\n", robot.Position.X, robot.Position.Y)

	return robot
}

type TurnCommand struct {
	angle int
}

func NewTurnCommand(angle int) TurnCommand {
	return TurnCommand{
		angle: angle,
	}
}

func (tc TurnCommand) Execute(r domain.Robot) domain.Robot {
	robot := r

	robot.Angle += tc.angle
	fmt.Printf("ANGLE %v\n", tc.angle)

	return robot
}

type SetCommand struct {
	state domain.State
}

func NewSetCommand(value domain.State) SetCommand {
	return SetCommand{
		state: value,
	}
}

func (sc SetCommand) Execute(r domain.Robot) domain.Robot {
	robot := r

	robot.State = sc.state
	fmt.Printf("STATE %v\n", sc.state)

	return robot
}

type StartCommand struct {
}

func NewStartCommand() StartCommand {
	return StartCommand{}
}

func (sc StartCommand) Execute(r domain.Robot) domain.Robot {
	fmt.Printf("START WITH %v\n", r.State)

	return r
}

type StopCommand struct {
}

func NewStopCommand() StopCommand {
	return StopCommand{}
}

func (sc StopCommand) Execute(r domain.Robot) domain.Robot {
	fmt.Printf("STOP\n")

	return r
}
