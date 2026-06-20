package main

import (
	"clean-architecture/command/internal/controller"
	"clean-architecture/command/internal/domain"
	"clean-architecture/command/internal/robo"
)

func main() {
	cmds := []domain.Commander{
		robo.NewMoveCommand(100),
		robo.NewTurnCommand(-90), 
		robo.NewSetCommand(domain.Soap),
		robo.NewStartCommand(),
		robo.NewMoveCommand(50),
		robo.NewStopCommand(),
	}
	
	controller := controller.New(cmds)
	controller.Run(domain.Robot{})
}
