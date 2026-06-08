package main

import (
	"net/http"
)

func main() {
	// robot := robo.New()

	// roboHandler := api.New(robot)

	mux := http.NewServeMux()
	// mux.HandleFunc("/robot/execute", roboHandler.Handler)
	http.ListenAndServe(":8080", mux)
}
