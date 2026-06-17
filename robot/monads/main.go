package main

import (
	"fmt"
	"strings"

	"clean-architecture/robot/monads/internal/domain"
	"clean-architecture/robot/monads/internal/monadrobot"
	"clean-architecture/robot/monads/internal/purerobot"
)

func main() {
	program :=
		monadrobot.Start().
			Then(monadrobot.Move(100)).
			Then(monadrobot.Turn(-90)).
			Then(monadrobot.SetState(domain.Soap)).
			Then(monadrobot.Move(50)).
			Then(monadrobot.Turn(90)).
			Then(monadrobot.SetState(domain.Brush)).
			Then(monadrobot.Move(30)).
			Then(monadrobot.Stop())

	log, robot := program.Run(purerobot.Robot{})

	fmt.Println("=== execution log ===")
	fmt.Println(strings.Join(log, "\n"))
	fmt.Println()
	fmt.Printf("=== final state ===\npos   (%d, %d)\nangle %d\nstate %s\n",
		robot.Position.X, robot.Position.Y, robot.Angle, robot.State)
}
