package main

import (
	"fmt"
	"net/http"
	"sync"

	"clean-architecture/robot/functional_di_v2/api"
	"clean-architecture/robot/functional_di_v2/internal/domain"
	"clean-architecture/robot/functional_di_v2/internal/purerobot"
)

func main() {
	var (
		state purerobot.Robot
		mu    sync.Mutex
	)

	execute := api.Execute(func(cmd domain.Cmd) {
		mu.Lock()
		defer mu.Unlock()

		state = purerobot.Apply(state, cmd)

		switch cmd.Name {
		case domain.Move:
			fmt.Printf("POS %v %v\n", state.Position.X, state.Position.Y)
		case domain.Turn:
			fmt.Printf("ANGLE %v\n", state.Angle)
		case domain.Set:
			fmt.Printf("STATE %v\n", state.State)
		case domain.Start:
			fmt.Printf("START WITH %v\n", state.State)
		case domain.Stop:
			fmt.Printf("STOP\n")
		}
	})

	handler := api.New(execute)

	mux := http.NewServeMux()
	mux.HandleFunc("/robot/execute", handler)
	http.ListenAndServe(":8080", mux)
}
