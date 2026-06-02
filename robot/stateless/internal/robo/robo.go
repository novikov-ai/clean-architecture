package robo

import (
	"clean-architecture/robot/stateless/internal/domain"
	"fmt"
	"strconv"
	"strings"
)

func Do(s domain.RoboState, commands ...string) domain.RoboState {
	for _, cmd := range commands {
		s = apply(s, strings.Split(cmd, " "))
	}
	return s
}

func apply(s domain.RoboState, args []string) domain.RoboState {
	if len(args) == 0 {
		return s
	}
	cmd, param := domain.Command(args[0]), args[1:]

	switch cmd {
	case domain.Start:
		fmt.Printf("START WITH %v\n", s.State)
	case domain.Stop:
		fmt.Printf("STOP\n")
	case domain.Move:
		s = move(s, intParam(param))
	case domain.Turn:
		s.Angle += intParam(param)
		fmt.Printf("ANGLE %v\n", s.Angle)
	case domain.Set:
		if len(param) > 0 {
			s.State = domain.State(param[0])
			fmt.Printf("STATE %v\n", s.State)
		}
	}
	return s
}

func move(s domain.RoboState, steps int) domain.RoboState {
	angle := ((s.Angle % 360) + 360) % 360
	switch angle {
	case 0:
		s.Position.X += steps
	case 90:
		s.Position.Y += steps
	case 180:
		s.Position.X -= steps
	case 270:
		s.Position.Y -= steps
	}
	fmt.Printf("POS %v %v\n", s.Position.X, s.Position.Y)
	return s
}

func intParam(args []string) int {
	if len(args) == 0 {
		return 0
	}
	n, _ := strconv.Atoi(args[0])
	return n
}
