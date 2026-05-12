package main

import "fmt"

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

func main() {
	r := New()
	r.Turn(-90)
	r.Set(soap)
	r.Start()
	r.Move(50)
	r.Stop()
}
