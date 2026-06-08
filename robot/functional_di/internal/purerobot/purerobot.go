package purerobot

import "clean-architecture/robot/functional_di/internal/domain"

type Robot struct {
	Position domain.Position
	Angle    int
	State    domain.State
}

func Move(r Robot, steps int) Robot {
	angle := ((r.Angle % 360) + 360) % 360

	switch angle {
	case 0:
		r.Position.X += steps
	case 90:
		r.Position.Y += steps
	case 180:
		r.Position.X -= steps
	case 270:
		r.Position.Y -= steps
	}

	return r
}

func Turn(r Robot, angle int) Robot {
	r.Angle += angle
	return r
}

func Set(r Robot, s domain.State) Robot {
	r.State = s
	return r
}

func Start(r Robot) Robot { return r }
func Stop(r Robot) Robot  { return r }
