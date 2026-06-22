package main

import (
	"clean-architecture/event_sourcing/internal/command_handler"
	"clean-architecture/event_sourcing/internal/event_store"
	"net/http"
)

func main() {
	eventStorage := event_store.New()
	
	cmdHandler := command_handler.New(eventStorage)
	
	mux := http.NewServeMux()
	mux.HandleFunc("/robot/command", cmdHandler.Handle)
	http.ListenAndServe(":8080", mux)
}
