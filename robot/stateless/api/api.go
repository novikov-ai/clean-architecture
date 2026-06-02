package api

import (
	"clean-architecture/robot/stateless/internal/domain"
	"clean-architecture/robot/stateless/internal/robo"
	"encoding/json"
	"net/http"
)

type request struct {
	State    domain.RoboState `json:"state"`
	Commands []string         `json:"commands"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	newState := robo.Do(req.State, req.Commands...)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newState)
}
