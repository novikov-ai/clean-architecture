package main

import (
	"clean-architecture/robot/stateless/api"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/robot/execute", api.Handler)
	http.ListenAndServe(":8080", mux)
}
