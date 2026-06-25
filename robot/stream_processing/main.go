package main

import (
	"net/http"

	"clean-architecture/robot/stream_processing/internal/command_handler"
	"clean-architecture/robot/stream_processing/internal/domain"
	"clean-architecture/robot/stream_processing/internal/event_store"
	"clean-architecture/robot/stream_processing/internal/processors"
	"clean-architecture/robot/stream_processing/internal/projector"
)

func main() {
	store := event_store.New()

	initial := domain.RobotState{X: 0, Y: 0, Angle: 0, State: int(domain.Water)}
	proj := projector.New(initial)

	processors.NewMovementProcessor(store, proj)
	processors.NewStateProcessor(store, proj)
	processors.NewLoggingProcessor(store)

	handler := command_handler.New("robot-1", store)

	mux := http.NewServeMux()
	mux.HandleFunc("/robot/command", handler.Handle)
	http.ListenAndServe(":8080", mux)
}
