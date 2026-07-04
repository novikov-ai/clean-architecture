package ast

import "clean-architecture/robot/ast/internal/domain"

type MoveResponse struct {
	Distance    int
	NewX, NewY  int
	Success     bool
}

type TurnResponse struct {
	Angle    int
	NewAngle int
	Success  bool
}

type StateResponse struct {
	OldState domain.State
	NewState domain.State
	Success  bool
}
