package monadrobot

import (
	"fmt"

	"clean-architecture/robot/monads/internal/domain"
	"clean-architecture/robot/monads/internal/purerobot"
)

type StateM[S, A any] struct {
	run func(S) (A, S)
}

func Bind[S, A, B any](m StateM[S, A], f func(A) StateM[S, B]) StateM[S, B] {
	return StateM[S, B]{
		run: func(s S) (B, S) {
			a, s1 := m.run(s)
			return f(a).run(s1)
		},
	}
}

type M struct {
	run func(purerobot.Robot) ([]string, purerobot.Robot)
}

func lift(sm StateM[purerobot.Robot, []string]) M {
	return M{run: sm.run}
}

func (m M) Then(next M) M {
	return lift(Bind(
		StateM[purerobot.Robot, []string]{run: m.run},
		func(log []string) StateM[purerobot.Robot, []string] {
			return StateM[purerobot.Robot, []string]{
				run: func(r purerobot.Robot) ([]string, purerobot.Robot) {
					log2, r2 := next.run(r)
					return append(log, log2...), r2
				},
			}
		},
	))
}

func (m M) Run(r purerobot.Robot) ([]string, purerobot.Robot) {
	return m.run(r)
}

func Move(steps int) M {
	return M{run: func(r purerobot.Robot) ([]string, purerobot.Robot) {
		r = purerobot.Move(r, steps)
		return []string{fmt.Sprintf("move(%d)  -> pos(%d, %d)", steps, r.Position.X, r.Position.Y)}, r
	}}
}

func Turn(angle int) M {
	return M{run: func(r purerobot.Robot) ([]string, purerobot.Robot) {
		r = purerobot.Turn(r, angle)
		return []string{fmt.Sprintf("turn(%d) -> angle(%d)", angle, r.Angle)}, r
	}}
}

func SetState(s domain.State) M {
	return M{run: func(r purerobot.Robot) ([]string, purerobot.Robot) {
		r = purerobot.Set(r, s)
		return []string{fmt.Sprintf("set(%s)", s)}, r
	}}
}

func Start() M {
	return M{run: func(r purerobot.Robot) ([]string, purerobot.Robot) {
		r = purerobot.Start(r)
		return []string{fmt.Sprintf("start    -> state(%s)", r.State)}, r
	}}
}

func Stop() M {
	return M{run: func(r purerobot.Robot) ([]string, purerobot.Robot) {
		r = purerobot.Stop(r)
		return []string{"stop"}, r
	}}
}
