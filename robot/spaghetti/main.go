package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Robot interface {
	Move(int)
	Turn(int)
	Set(state)
	Start()
	Stop()
}

type position struct {
	x int
	y int
}

type command string

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

func (r *Robo) Do(cmd string) {
	actions := strings.Split(cmd, " ")
	if len(actions) == 0 {
		return
	}

	switch command(actions[0]) {
	case move:
		if len(actions) != 2 {
			return
		}

		value, _ := strconv.Atoi(actions[1])
		r.Move(value)
	case turn:
		if len(actions) != 2 {
			return
		}

		value, _ := strconv.Atoi(actions[1])
		r.Turn(value)
	case set:
		if len(actions) != 2 {
			return
		}

		r.Set(state(actions[1]))
	case start:
		r.Start()
	case stop:
		r.Stop()
	}
}

func main() {
	r := New()
	for _, c := range []string{"move 100", "turn -90", "set soap", "start", "move 50", "stop"} {
		r.Do(c)
	}
}
