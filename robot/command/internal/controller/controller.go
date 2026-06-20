package controller

import (
	"clean-architecture/command/internal/domain"
)

type controller struct {
	commands []domain.Commander
}

func New(cmds []domain.Commander) controller {
	return controller{
		commands: cmds,
	}
}

func (c controller) Run(r domain.Robot) {
	var robot domain.Robot = r
	for _, cmd := range c.commands {
		robot = cmd.Execute(robot)
	}
}
