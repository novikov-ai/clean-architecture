package purerobot

import "clean-architecture/robot/monads_v3/internal/domain"

const (
	areaMin = 0
	areaMax = 100
)

type Robot struct {
	position domain.Position
	angle    int
	state    domain.State
	water    int
	soap     int
	stuck    bool
}

func (r Robot) Position() domain.Position { return r.position }
func (r Robot) Angle() int                { return r.angle }
func (r Robot) State() domain.State       { return r.state }
func (r Robot) Water() int                { return r.water }
func (r Robot) Soap() int                 { return r.soap }

type Capabilities struct {
	Move     func(steps int) (Robot, Capabilities)
	Turn     func(angle int) (Robot, Capabilities)
	SetWater func() (Robot, Capabilities)
	SetSoap  func() (Robot, Capabilities)
	SetBrush func() (Robot, Capabilities)
	Stop     func() (Robot, Capabilities)
	Start    func() (Robot, Capabilities)
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

func normalizeAngle(a int) int {
	return ((a % 360) + 360) % 360
}

func New(water, soap int) (Robot, Capabilities) {
	r := Robot{water: water, soap: soap}
	return r, stoppedCapabilities(r)
}

func stoppedCapabilities(r Robot) Capabilities {
	return Capabilities{
		Start: func() (Robot, Capabilities) {
			return r, runningCapabilities(r)
		},
	}
}

func runningCapabilities(r Robot) Capabilities {
	caps := Capabilities{
		Turn: func(angle int) (Robot, Capabilities) {
			nr := r
			nr.angle = normalizeAngle(nr.angle + angle)
			nr.stuck = false
			return nr, runningCapabilities(nr)
		},
		SetBrush: func() (Robot, Capabilities) {
			nr := r
			nr.state = domain.Brush
			return nr, runningCapabilities(nr)
		},
		Stop: func() (Robot, Capabilities) {
			return r, stoppedCapabilities(r)
		},
	}

	if !r.stuck {
		caps.Move = func(steps int) (Robot, Capabilities) {
			nr := move(r, steps)
			return nr, runningCapabilities(nr)
		}
	}

	if r.water > 0 {
		caps.SetWater = func() (Robot, Capabilities) {
			nr := r
			nr.water--
			nr.state = domain.Water
			return nr, runningCapabilities(nr)
		}
	}
	if r.soap > 0 {
		caps.SetSoap = func() (Robot, Capabilities) {
			nr := r
			nr.soap--
			nr.state = domain.Soap
			return nr, runningCapabilities(nr)
		}
	}

	return caps
}

func move(r Robot, steps int) Robot {
	x, y := r.position.X, r.position.Y
	switch r.angle {
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
	r.position = domain.Position{X: cx, Y: cy}
	r.stuck = cx != x || cy != y
	return r
}
