package main

import (
	"fmt"
	"net/http"
	"sync"

	"clean-architecture/robot/functional_di/api"
	"clean-architecture/robot/functional_di/internal/domain"
	"clean-architecture/robot/functional_di/internal/purerobot"
)

func main() {
	var (
		state purerobot.Robot
		mu    sync.Mutex
	)

	move := domain.ActionCommand(func(steps int) {
		mu.Lock()
		defer mu.Unlock()
		state = purerobot.Move(state, steps)
		fmt.Printf("POS %v %v\n", state.Position.X, state.Position.Y)
	})

	turn := domain.ActionCommand(func(angle int) {
		mu.Lock()
		defer mu.Unlock()
		state = purerobot.Turn(state, angle)
		fmt.Printf("ANGLE %v\n", state.Angle)
	})

	set := domain.StateCommand(func(s domain.State) {
		mu.Lock()
		defer mu.Unlock()
		state = purerobot.Set(state, s)
		fmt.Printf("STATE %v\n", state.State)
	})

	start := domain.ModeCommand(func() {
		mu.Lock()
		defer mu.Unlock()
		state = purerobot.Start(state)
		fmt.Printf("START WITH %v\n", state.State)
	})

	stop := domain.ModeCommand(func() {
		mu.Lock()
		defer mu.Unlock()
		state = purerobot.Stop(state)
		fmt.Printf("STOP\n")
	})

	handler := api.New(move, turn, set, start, stop)

	mux := http.NewServeMux()
	mux.HandleFunc("/robot/execute", handler)
	http.ListenAndServe(":8080", mux)
}
