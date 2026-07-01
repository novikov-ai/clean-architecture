package monadrobot

import (
	"fmt"

	"clean-architecture/robot/monads_v2/internal/domain"
	"clean-architecture/robot/monads_v2/internal/purerobot"
)

type M[A any] struct {
	run func(purerobot.Robot) (A, []string, purerobot.Robot)
}

func Bind[A, B any](m M[A], f func(A) M[B]) M[B] {
	return M[B]{
		run: func(r purerobot.Robot) (B, []string, purerobot.Robot) {
			a, log1, r1 := m.run(r)
			b, log2, r2 := f(a).run(r1)
			return b, append(log1, log2...), r2
		},
	}
}

func Then[A, B any](m M[A], next M[B]) M[B] {
	return Bind(m, func(A) M[B] { return next })
}

func (m M[A]) Run(r purerobot.Robot) (A, []string, purerobot.Robot) {
	return m.run(r)
}

func Move(steps int) M[domain.MoveResult] {
	return M[domain.MoveResult]{run: func(r purerobot.Robot) (domain.MoveResult, []string, purerobot.Robot) {
		newR, res := purerobot.Move(r, steps)
		log := fmt.Sprintf("move(%d)  -> pos(%d, %d) [%s]", steps, newR.Position.X, newR.Position.Y, res)
		return res, []string{log}, newR
	}}
}

func Turn(angle int) M[struct{}] {
	return M[struct{}]{run: func(r purerobot.Robot) (struct{}, []string, purerobot.Robot) {
		r = purerobot.Turn(r, angle)
		return struct{}{}, []string{fmt.Sprintf("turn(%d) -> angle(%d)", angle, r.Angle)}, r
	}}
}

func SetState(s domain.State) M[domain.SetStateResult] {
	return M[domain.SetStateResult]{run: func(r purerobot.Robot) (domain.SetStateResult, []string, purerobot.Robot) {
		newR, res := purerobot.Set(r, s)
		log := fmt.Sprintf("set(%s) -> %s", s, res)
		return res, []string{log}, newR
	}}
}

func Start() M[struct{}] {
	return M[struct{}]{run: func(r purerobot.Robot) (struct{}, []string, purerobot.Robot) {
		r = purerobot.Start(r)
		return struct{}{}, []string{fmt.Sprintf("start    -> state(%s)", r.State)}, r
	}}
}

func Stop() M[struct{}] {
	return M[struct{}]{run: func(r purerobot.Robot) (struct{}, []string, purerobot.Robot) {
		r = purerobot.Stop(r)
		return struct{}{}, []string{"stop"}, r
	}}
}
