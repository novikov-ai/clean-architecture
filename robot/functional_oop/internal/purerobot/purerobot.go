package purerobot

import "clean-architecture/functional_oop/internal/domain"

type Robot struct {
	position domain.Position
	angle    int
	state    domain.State
}

func New() Robot {
	return Robot{}
}

func Move(r Robot, steps int) Robot {
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

	return r
}

func Turn(r Robot, angle int) Robot {
	r.angle += angle
	return r
}

func Set(r Robot, s domain.State) Robot {
	r.state = s
	return r
}

func Start(r Robot) Robot { return r }
func Stop(r Robot) Robot  { return r }

func (r Robot) ShowPosition() domain.Position {
	return r.position
}

func (r Robot) ShowAngle() int {
	return r.angle
}

func (r Robot) ShowState() domain.State {
	return r.state
}
