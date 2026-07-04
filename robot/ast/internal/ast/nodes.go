package ast

import (
	"clean-architecture/robot/ast/internal/domain"
	"fmt"
)

type StopNode struct{}

func Stop() StopNode { return StopNode{} }

func (n StopNode) Interpret(robot domain.Robot) domain.Robot {
	fmt.Println("[Stop]")
	return robot
}

type MoveNode struct {
	Distance int
	Next     func(MoveResponse) Node
}

func Move(distance int, next func(MoveResponse) Node) MoveNode {
	return MoveNode{Distance: distance, Next: next}
}

func (n MoveNode) Interpret(robot domain.Robot) domain.Robot {
	angle := ((robot.Angle % 360) + 360) % 360
	newRobot := robot
	switch angle {
	case 0:
		newRobot.Position.X += n.Distance
	case 90:
		newRobot.Position.Y += n.Distance
	case 180:
		newRobot.Position.X -= n.Distance
	case 270:
		newRobot.Position.Y -= n.Distance
	}
	fmt.Printf("[Move %d] pos=(%d, %d)\n", n.Distance, newRobot.Position.X, newRobot.Position.Y)

	resp := MoveResponse{
		Distance: n.Distance,
		NewX:     newRobot.Position.X,
		NewY:     newRobot.Position.Y,
		Success:  true,
	}
	return n.Next(resp).Interpret(newRobot)
}

type TurnNode struct {
	Angle int
	Next  func(TurnResponse) Node
}

func Turn(angle int, next func(TurnResponse) Node) TurnNode {
	return TurnNode{Angle: angle, Next: next}
}

func (n TurnNode) Interpret(robot domain.Robot) domain.Robot {
	newRobot := robot
	newRobot.Angle += n.Angle
	fmt.Printf("[Turn %d] angle=%d\n", n.Angle, newRobot.Angle)

	resp := TurnResponse{
		Angle:    n.Angle,
		NewAngle: newRobot.Angle,
		Success:  true,
	}
	return n.Next(resp).Interpret(newRobot)
}

type SetStateNode struct {
	State domain.State
	Next  func(StateResponse) Node
}

func SetState(state domain.State, next func(StateResponse) Node) SetStateNode {
	return SetStateNode{State: state, Next: next}
}

func (n SetStateNode) Interpret(robot domain.Robot) domain.Robot {
	newRobot := robot
	old := newRobot.State
	newRobot.State = n.State
	fmt.Printf("[SetState %s → %s]\n", old, n.State)

	resp := StateResponse{
		OldState: old,
		NewState: n.State,
		Success:  true,
	}
	return n.Next(resp).Interpret(newRobot)
}
