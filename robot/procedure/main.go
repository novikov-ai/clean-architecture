package main

import (
	"fmt"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
}

type command string

type (
	ModeCommand   func()
	ActionCommand func(int)
	StateCommand  func(state)
)

const (
	move  command = "move"
	turn  command = "turn"
	set   command = "set"
	start command = "start"
	stop  command = "stop"
)

type state string

const (
	water state = "water"
	soap  state = "soap"
	brush state = "brush"
)

var (
	angle         = 0
	robotPosition = position{0, 0}
	robotState    = water
)

func Move(steps int) {
	angle := ((angle % 360) + 360) % 360

	switch angle {
	case 0:
		robotPosition.x += steps
	case 90:
		robotPosition.y += steps
	case 180:
		robotPosition.x -= steps
	case 270:
		robotPosition.y -= steps
	}

	fmt.Printf("POS %v %v\n", robotPosition.x, robotPosition.y)
}

func Turn(angle int) {
	angle += angle
	fmt.Printf("ANGLE %v\n", angle)
}

func Set(s state) {
	robotState = s
	fmt.Printf("STATE %v\n", s)
}

func Start() {
	fmt.Printf("START WITH %v\n", robotState)
}

func Stop() {
	fmt.Printf("STOP\n")
}

func Do(commands ...string) {
	for _, cmd := range commands {
		splitted := strings.Split(cmd, " ")
		parseModeCommand(splitted)()

		action, arg := parseActionCommand(splitted)
		action(arg)

		command, state := parseStateCommand(splitted)
		command(state)
	}
}

func parseModeCommand(args []string) ModeCommand {
	if len(args) != 1 {
		return func() {}
	}

	switch command(args[0]) {
	case start:
		return Start
	case stop:
		return Stop
	default:
		return func() {}
	}
}

func parseActionCommand(args []string) (ActionCommand, int) {
	if len(args) != 2 {
		return func(int) {}, 0
	}

	value, _ := strconv.Atoi(args[1])

	switch command(args[0]) {
	case move:
		return Move, value
	case turn:
		return Turn, value
	default:
		return func(i int) {}, 0
	}
}

func parseStateCommand(args []string) (StateCommand, state) {
	if len(args) != 2 {
		return func(state) {}, ""
	}

	switch command(args[0]) {
	case set:
		return Set, state(args[1])
	default:
		return func(s state) {}, ""
	}
}

func main() {
	commands := []string{"move 100", "turn -90", "set soap", "start", "move 50", "stop"}

	Do(commands...)
}
