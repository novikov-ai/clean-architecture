package controller

import (
	"clean-architecture/robot/ast/internal/ast"
	"clean-architecture/robot/ast/internal/domain"
)

func Run(program ast.Node, initial domain.Robot) domain.Robot {
	return program.Interpret(initial)
}
