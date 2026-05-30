package api

import (
	"clean-architecture/robot/hipster/internal/robo"
	"io"
	"net/http"
	"strings"
)

var robot = robo.New()

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	commands := strings.Split(strings.TrimSpace(string(body)), "\n")
	robot.Do(commands...)
}
