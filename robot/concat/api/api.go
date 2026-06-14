package api

import (
	"clean-architecture/robot/concat/internal/robo"
	"io"
	"net/http"
)

var robot = robo.New()

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	robot.Do(string(body))
}
