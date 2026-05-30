package robo

import (
	"clean-architecture/robot/module/domain"
	"fmt"
	"strconv"
	"strings"
)

type Robo struct {
	position domain.Position
	angle    int
	state    domain.State
}

func New() *Robo {
	return &Robo{}
}

func (r *Robo) Move(steps int) {
	angle := ((r.angle % 360) + 360) % 360

	switch angle {
	case 0:
		r.position.X += steps
	case 90:
		r.position.Y += steps
	case 180:
		r.position.X -= steps
	case 270:
		r.position.Y -= steps
	}

	fmt.Printf("POS %v %v\n", r.position.X, r.position.Y)
}

func (r *Robo) Turn(angle int) {
	r.angle += angle
	fmt.Printf("ANGLE %v\n", r.angle)
}

func (r *Robo) Set(s domain.State) {
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

func (r *Robo) parseModeCommand(args []string) domain.ModeCommand {
	if len(args) != 1 {
		return func() {}
	}

	switch domain.Command(args[0]) {
	case domain.Start:
		return r.Start
	case domain.Stop:
		return r.Stop
	default:
		return func() {}
	}
}

func (r *Robo) parseActionCommand(args []string) (domain.ActionCommand, int) {
	if len(args) != 2 {
		return func(int) {}, 0
	}

	value, _ := strconv.Atoi(args[1])

	switch domain.Command(args[0]) {
	case domain.Move:
		return r.Move, value
	case domain.Turn:
		return r.Turn, value
	default:
		return func(i int) {}, 0
	}
}

func (r *Robo) parseStateCommand(args []string) (domain.StateCommand, domain.State) {
	if len(args) != 2 {
		return func(domain.State) {}, ""
	}

	switch domain.Command(args[0]) {
	case domain.Set:
		return r.Set, domain.State(args[1])
	default:
		return func(s domain.State) {}, ""
	}
}
