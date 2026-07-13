package main

import (
	"fmt"

	"clean-architecture/robot/monads_v3/internal/purerobot"
)

func report(name string, r purerobot.Robot) {
	fmt.Printf("=== %s ===\n", name)
	fmt.Printf("final: pos(%d, %d) angle(%d) state(%s) water(%d) soap(%d)\n\n",
		r.Position().X, r.Position().Y, r.Angle(), r.State(), r.Water(), r.Soap())
}

func main() {
	r, caps := purerobot.New(1, 1)
	r, caps = caps.Start()
	r, caps = caps.Move(50)
	r, caps = caps.Turn(90)
	r, caps = caps.SetSoap()
	r, caps = caps.Move(30)
	r, caps = caps.Stop()
	report("happy path", r)

	r, caps = purerobot.New(1, 1)
	r, caps = caps.Start()
	r, caps = caps.Move(150) // clamped at the wall
	if caps.Move == nil {
		fmt.Println("hits barrier: Move capability is gone, further attempts are impossible")
	}
	r, caps = caps.Turn(90)
	if caps.Move != nil {
		fmt.Println("after turn: Move capability is back")
	}
	r, caps = caps.Stop()
	report("hits barrier", r)

	r, caps = purerobot.New(1, 0)
	r, caps = caps.Start()
	r, caps = caps.SetWater() // consumes the only unit of water
	if caps.SetWater == nil {
		fmt.Println("out of water: SetWater capability is gone, no more wasted attempts")
	}
	r, caps = caps.Stop()
	report("out of water", r)
}
