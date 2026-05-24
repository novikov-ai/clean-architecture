package main

import (
	"clean-architecture/robot/module/robo"
)

func main() {
	commands := []string{"move 100", "turn -90", "set soap", "start", "move 50", "stop"}

	r := robo.New()
	r.Do(commands...)
}
