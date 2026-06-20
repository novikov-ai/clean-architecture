package main

import (
	"clean-architecture/robot/command/controller"
	"clean-architecture/robot/command/internal/domain"
	"clean-architecture/robot/command/internal/robot"
)

func main() {
	cmds := []domain.Commander{
		MoveCommand{}, // 100
		TurnCommand{}, // -90
		SetCommand{}, // soap
		StartCommand{},
		MoveCommand{}, // 50
		StopCommand{},
	}
	
	controller := conroller.New(cmds)
	controller.Run()
}
