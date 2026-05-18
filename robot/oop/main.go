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

type Robo struct {
	position position
	angle    int
	state    state
}

func New() *Robo {
	return &Robo{}
}

func (r *Robo) Move(steps int) {
	angle := ((r.angle % 360) + 360) % 360

	switch angle {
	case 0:
		r.position.x += steps
	case 90:
		r.position.y += steps
	case 180:
		r.position.x -= steps
	case 270:
		r.position.y -= steps
	}

	fmt.Printf("POS %v %v\n", r.position.x, r.position.y)
}

func (r *Robo) Turn(angle int) {
	r.angle += angle
	fmt.Printf("ANGLE %v\n", r.angle)
}

func (r *Robo) Set(s state) {
	r.state = s
	fmt.Printf("STATE %v\n", s)
}

func (r *Robo) Start() {
	fmt.Printf("START WITH %v\n", r.state)
}

func (r *Robo) Stop() {
	fmt.Printf("STOP\n")
}

func (r *Robo) Do(commands ...string) {
	for _, cmd := range commands {
		splitted := strings.Split(cmd, " ")
		r.parseModeCommand(splitted)()

		action, arg := r.parseActionCommand(splitted)
		action(arg)

		command, state := r.parseStateCommand(splitted)
		command(state)
	}
}

func (r *Robo) parseModeCommand(args []string) ModeCommand {
	if len(args) != 1 {
		return func() {}
	}

	switch command(args[0]) {
	case start:
		return r.Start
	case stop:
		return r.Stop
	default:
		return func() {}
	}
}

func (r *Robo) parseActionCommand(args []string) (ActionCommand, int) {
	if len(args) != 2 {
		return func(int) {}, 0
	}

	value, _ := strconv.Atoi(args[1])

	switch command(args[0]) {
	case move:
		return r.Move, value
	case turn:
		return r.Turn, value
	default:
		return func(i int) {}, 0
	}
}

func (r *Robo) parseStateCommand(args []string) (StateCommand, state) {
	if len(args) != 2 {
		return func(state) {}, ""
	}

	switch command(args[0]) {
	case set:
		return r.Set, state(args[1])
	default:
		return func(s state) {}, ""
	}
}

func main() {
	commands := []string{"move 100", "turn -90", "set soap", "start", "move 50", "stop"}

	r := New()
	r.Do(commands...)
}
