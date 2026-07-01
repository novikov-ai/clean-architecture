package purerobot

import "clean-architecture/robot/monads_v2/internal/domain"

const (
	areaMin = 0
	areaMax = 100
)

type Robot struct {
	Position domain.Position
	Angle    int
	State    domain.State
	Water    int
	Soap     int
}

func clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func Move(r Robot, steps int) (Robot, domain.MoveResult) {
	angle := ((r.Angle % 360) + 360) % 360

	x, y := r.Position.X, r.Position.Y
	switch angle {
	case 0:
		x += steps
	case 90:
		y += steps
	case 180:
		x -= steps
	case 270:
		y -= steps
	}

	cx, cy := clamp(x, areaMin, areaMax), clamp(y, areaMin, areaMax)
	r.Position = domain.Position{X: cx, Y: cy}

	if cx != x || cy != y {
		return r, domain.HitBarrier
	}
	return r, domain.MovedOk
}

func Turn(r Robot, angle int) Robot {
	r.Angle += angle
	return r
}

func Set(r Robot, s domain.State) (Robot, domain.SetStateResult) {
	switch s {
	case domain.Water:
		if r.Water <= 0 {
			return r, domain.NoWater
		}
		r.Water--
	case domain.Soap:
		if r.Soap <= 0 {
			return r, domain.NoSoap
		}
		r.Soap--
	}

	r.State = s
	return r, domain.SetStateOk
}

func Start(r Robot) Robot { return r }
func Stop(r Robot) Robot  { return r }
