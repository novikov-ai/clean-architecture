package ast

import "clean-architecture/robot/ast/internal/domain"

type Node interface {
	Interpret(robot domain.Robot) domain.Robot
}
