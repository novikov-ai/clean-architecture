package main

import (
	"fmt"
	"strings"

	"clean-architecture/robot/monads_v2/internal/domain"
	"clean-architecture/robot/monads_v2/internal/monadrobot"
	"clean-architecture/robot/monads_v2/internal/purerobot"
)

func moveOrStop(steps int, onOk monadrobot.M[struct{}]) monadrobot.M[struct{}] {
	return monadrobot.Bind(monadrobot.Move(steps), func(res domain.MoveResult) monadrobot.M[struct{}] {
		if res == domain.HitBarrier {
			return monadrobot.Stop()
		}
		return onOk
	})
}

func setStateOrStop(s domain.State, onOk monadrobot.M[struct{}]) monadrobot.M[struct{}] {
	return monadrobot.Bind(monadrobot.SetState(s), func(res domain.SetStateResult) monadrobot.M[struct{}] {
		if res != domain.SetStateOk {
			return monadrobot.Stop()
		}
		return onOk
	})
}

func run(name string, r purerobot.Robot, program monadrobot.M[struct{}]) {
	_, log, final := program.Run(r)

	fmt.Printf("=== %s ===\n", name)
	fmt.Println(strings.Join(log, "\n"))
	fmt.Printf("final: pos(%d, %d) angle(%d) state(%s) water(%d) soap(%d)\n\n",
		final.Position.X, final.Position.Y, final.Angle, final.State, final.Water, final.Soap)
}

func main() {
	happyPath := monadrobot.Then(monadrobot.Start(),
		moveOrStop(50,
			monadrobot.Then(monadrobot.Turn(90),
				setStateOrStop(domain.Soap,
					moveOrStop(30, monadrobot.Stop()),
				),
			),
		),
	)
	run("happy path", purerobot.Robot{Water: 1, Soap: 1}, happyPath)

	hitsBarrier := monadrobot.Then(monadrobot.Start(),
		moveOrStop(150,
			setStateOrStop(domain.Soap, monadrobot.Stop()),
		),
	)
	run("hits barrier", purerobot.Robot{Water: 1, Soap: 1}, hitsBarrier)

	outOfSoap := monadrobot.Then(monadrobot.Start(),
		moveOrStop(50,
			setStateOrStop(domain.Soap,
				moveOrStop(30, monadrobot.Stop()),
			),
		),
	)
	run("out of soap", purerobot.Robot{Water: 1, Soap: 0}, outOfSoap)
}
