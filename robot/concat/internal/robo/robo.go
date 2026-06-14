package robo

import (
	"clean-architecture/robot/concat/internal/domain"
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

func (r *Robo) Do(commands string) {
	splited := strings.Split(commands, " ")

	var value *int

	for _, cmd := range splited {
		switch cmd {
		case string(domain.Start):
			r.Start()
		case string(domain.Stop):
			r.Stop()
		case string(domain.Set):
			continue
		default:
			if isState(cmd) {
				r.Set(cmd)
				continue
			}

			switch cmd {
			case string(domain.Turn):
				if value != nil {
					r.Turn(*value)
					value = nil
					continue
				}
			case string(domain.Move):
				if value != nil {
					r.Move(*value)
					value = nil
					continue
				}
			}

			v, _ := strconv.Atoi(cmd)
			value = &v
		}
	}
}

func isState(cmd string) bool {
	switch cmd {
	case string(domain.Brush), string(domain.Soap), string(domain.Water):
		return true
	}

	return false
}
